package utilsgo

import (
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Filename  string
	Path      string
	Extension string
	IsDir     bool
	Size      int64
	Filemode  os.FileMode
}

func IsFileOrFolderExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		return err
	}
	return nil
}

func DeleteDirectory(path string) error {
	err := os.RemoveAll(path)
	return err
}

func ReadFilesFromDir(path string) (list []File, err2 error) {
	err2 = filepath.Walk(path,
		func(fpath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			ext := strings.Split(info.Name(), ".")[len(strings.Split(info.Name(), "."))-1]
			list = append(list, File{
				Filename:  info.Name(),
				Extension: ext,
				Path:      fpath,
				IsDir:     info.IsDir(),
				Size:      info.Size(),
				Filemode:  info.Mode(),
			})
			return nil
		})
	return
}
