package utilsgo

import (
	"os"
)

func IsFileOrFolderExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		return err
	}
	return nil
}

func DeleteDirectory(path string) error {
	err := os.RemoveAll(path)
	return err
}
