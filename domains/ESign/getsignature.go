package ESign

import (
	"archive/zip"
	"io"
	"net/http"
	"os"

	"DRD/db"

	"github.com/gin-gonic/gin"
)

func getSignature(c *gin.Context) {
	// Input Tipe Text
	userId := c.PostForm("UserId")
	var user db.User
	dbInstance := db.GetDb()
	dbInstance.First(&user, "ID = ?", userId)
	if user.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	folderPath := "./Images/Users/" + user.ID + "/"
	if user.Signature != "" {
		c.Writer.Header().Set("Content-type", "application/octet-stream")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename=Signature_"+user.ID+".zip")
	} else {
		c.JSON(http.StatusAccepted, UserSignatureStatus{UserId: user.ID, SignStatus: 0})
		return
	}
	archiveSign := zip.NewWriter(c.Writer)

	signature, _ := os.Open(folderPath + user.Signature)
	sign, _ := archiveSign.Create(user.Signature)

	io.Copy(sign, signature)
	// Optional
	if user.Signature2 != "" {
		signature2, _ := os.Open(folderPath + user.Signature2)
		sign2, _ := archiveSign.Create(user.Signature2)
		io.Copy(sign2, signature2)
	}
	if user.Signature3 != "" {
		signature3, _ := os.Open(folderPath + user.Signature3)
		sign3, _ := archiveSign.Create(user.Signature3)
		io.Copy(sign3, signature3)
	}
	archiveSign.Close()
	c.JSON(200, gin.H{})
}
