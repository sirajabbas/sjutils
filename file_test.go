package utilsgo

import (
	"fmt"
	"log"
	"testing"
)

func TestIsFileOrFolderExists(t *testing.T) {
	filename := "./go.mod"
	fileExists := IsFileOrFolderExists(filename)
	t.Log(fmt.Sprintf("file: %s exists:%v", filename, fileExists))

	path := "/home/testpathdoesnotexists"
	folderExists := IsFileOrFolderExists(path)
	t.Log(fmt.Sprintf("folder: %s exists:%v", path, folderExists))
}

func TestCreateDirectory(t *testing.T) {
	path := "./testdir"
	err := CreateDirectory(path)
	if err != nil {
		t.Fatal("CreateDirectory failed: ", err)
	}
}

func TestDeleteDirectory(t *testing.T) {
	path := "./testdir"
	err := DeleteDirectory(path)
	if err != nil {
		log.Fatal("Error executing DeleteDirectory", err)
	}
}
