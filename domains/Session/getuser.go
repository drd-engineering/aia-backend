package session

import (
	"fmt"
	"time"

	"DRD/db"

	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) {
	token := c.Param("token")
	fmt.Println(token)
	dbInstance := db.GetDb()
	var session db.Session
	dbInstance.First(&session, "token = ?", token)
	if session.ID <= 0 {
		c.AbortWithStatus(404)
	}
	var user db.User
	dbInstance.First(&user, "id = ?", session.UserID)
	if user.ID == "" {
		c.AbortWithStatus(404)
	}

	userReturn := User{
		KtpNumber:    user.KtpNumber,
		Name:         user.Name,
		Address:      user.Address,
		Gender:       user.Gender,
		PlaceOfBirth: user.PlaceOfBirth,
		Cityzenship:  user.Cityzenship,
		DateOfBirth:  "",
	}
	if session.CreatedAt.Add(time.Duration(session.Duration) * time.Second).Before(time.Now()) {
		linkStatus := LinkStatus{
			Token:      session.Token,
			LinkStatus: 0,
		}
		c.JSON(200, linkStatus)
		return
	}
	linkStatus := LinkStatus{
		UserId:     user.ID,
		User:       userReturn,
		Token:      session.Token,
		LinkStatus: 1,
	}
	c.JSON(200, linkStatus)
}
