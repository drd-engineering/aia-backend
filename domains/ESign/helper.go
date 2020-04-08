package ESign

import (
	"os"
)

func ensureUserImageFolderExist(userId string) error {
	basePath := "Images/"
	_, err := os.Stat(basePath)
	if os.IsNotExist(err) {
		err = os.Mkdir(basePath, os.ModeDir)
	}
	if err != nil {
		return err
	}
	userPath := basePath + "Users/"
	_, err = os.Stat(userPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(userPath, os.ModeDir)
	}
	if err != nil {
		return err
	}
	userSpecificPath := userPath + userId + "/"
	_, err = os.Stat(userSpecificPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(userSpecificPath, os.ModeDir)
	}
	if err != nil {
		return err
	}
	return nil
}
