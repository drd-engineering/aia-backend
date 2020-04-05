package ESign

import (
	"net/http"
	"path/filepath"

	"DRD/db"

	"github.com/gin-gonic/gin"
)

func saveSigns(c *gin.Context) {
	// Input Tipe Text
	userId := c.PostForm("UserId")
	// Input Tipe File
	var user db.User
	dbInstance := db.GetDb()
	dbInstance.First(&user, "ID = ?", userId)
	if user.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	file, err := c.FormFile("Signature1")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "required file upload"})
		return
	}

	// Set Folder untuk menyimpan filenya
	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	dbInstance.Model(&user).Update("Signature", filename)

	// Adding Optional Signature
	file, err = c.FormFile("Signature2")
	if err == nil {
		// Set Folder untuk menyimpan filenya
		filename2 := filepath.Base(file.Filename)
		if err = c.SaveUploadedFile(file, filename2); err != nil {
			filename2 = ""
		}
		dbInstance.Model(&user).Update("Signature2", filename2)
	}
	file, err = c.FormFile("Signature3")
	if err == nil {
		// Set Folder untuk menyimpan filenya
		filename3 := filepath.Base(file.Filename)
		if err = c.SaveUploadedFile(file, filename3); err != nil {
			filename3 = ""
		}
		dbInstance.Model(&user).Update("Signature3", filename3)
	}
	// Response
	response := UserSignature{
		UserId: user.ID,
		Signature: []string{
			user.Signature,
			user.Signature2,
			user.Signature3,
		},
	}
	c.JSON(200, response)
}
