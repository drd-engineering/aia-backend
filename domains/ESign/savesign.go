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
	err = ensureUserImageFolderExist(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	folderPath := "./Images/Users/" + user.ID + "/"

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, folderPath+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	dbInstance.Model(&user).Update("Signature", filename)

	// Adding Optional Signature
	file, err = c.FormFile("Signature2")
	if err == nil {
		// Set Folder untuk menyimpan filenya
		filename = filepath.Base(file.Filename)
		err = c.SaveUploadedFile(file, folderPath+filename)
	}
	if err == nil {
		dbInstance.Model(&user).Update("Signature2", filename)
	}
	file, err = c.FormFile("Signature3")
	if err == nil {
		// Set Folder untuk menyimpan filenya
		filename = filepath.Base(file.Filename)
		err = c.SaveUploadedFile(file, folderPath+filename)
	}
	if err == nil {
		dbInstance.Model(&user).Update("Signature3", filename)
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
