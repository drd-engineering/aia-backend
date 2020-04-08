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

	var dbInstance = db.GetDb()
	if input.KtpNumber == 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	userBirthDate, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	var userDb = db.User{
		KtpNumber:    input.KtpNumber,
		Name:         input.Name,
		Address:      input.Address,
		Gender:       input.Gender,
		PlaceOfBirth: input.PlaceOfBirth,
		Cityzenship:  input.Cityzenship,
		DateOfBirth:  userBirthDate,
	}

	// Disini panggil third party check dukcapil
	// If oke

	// Update Or Create
	dbInstance.Where(db.User{KtpNumber: userDb.KtpNumber}).Assign(&userDb).FirstOrCreate(&userDb)

	if userDb.ID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var sessionUser = db.Session{
		UserID:   userDb.ID,
		User:     userDb,
		Duration: 1800,
	}
	dbInstance.Order("created_at desc").First(&sessionUser)
	if sessionUser.ID != 0 {
		dbInstance.Model(&db.Session{ID: sessionUser.ID}).Update("duration", 0)
	}

	var newSessionUser = db.Session{
		User:     userDb,
		Duration: 1800,
	}

	dbInstance.Create(&newSessionUser)

	sessionLink := SessionLink{
		UserId: newSessionUser.User.ID,
		Path:   "http://linkbagasttd/" + newSessionUser.Token,
		Valid:  newSessionUser.CreatedAt.Add(time.Duration(newSessionUser.Duration) * time.Second),
	}

	c.JSON(200, sessionLink)
	// http://domain:3000
	// If the KTP is not valid with dukcapil's data
	// c.JSON(http.StatusInternalServerError, gin.H{})
	// return

}
