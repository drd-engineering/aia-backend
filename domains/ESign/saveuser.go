package ESign

import (
	"fmt"
	"net/http"
	"time"

	"DRD/db"

	"github.com/gin-gonic/gin"
)

func saveUser(c *gin.Context) {
	fmt.Println("Generating link")
	var input User
	c.ShouldBindJSON(&input)

	// Disini panggil third party check dukcapil
	// If oke

	var dbInstance = db.GetDb()
	var userDb = db.User{
		KtpNumber:    input.KtpNumber,
		Name:         input.Name,
		Address:      input.Address,
		Gender:       input.Gender,
		PlaceOfBirth: input.PlaceOfBirth,
		Cityzenship:  input.Cityzenship,
	}

	//kurang di date of birth
	fmt.Println(userDb)
	dbInstance.Create(&userDb)

	if userDb.ID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var sessionUser = db.Session{
		User:     userDb,
		Duration: 1800,
	}
	dbInstance.Create(&sessionUser)

	//Kurang di valid value
	sessionLink := SessionLink{
		UserId: sessionUser.User.ID,
		Path:   "http://linkbagasttd/" + sessionUser.Token,
		Valid:  sessionUser.CreatedAt.Add(time.Duration(sessionUser.Duration) * time.Second),
	}

	c.JSON(200, sessionLink)
	// http://domain:3000
	// If the KTP is not valid with dukcapil's data
	// c.JSON(http.StatusInternalServerError, gin.H{})
	// return
}
