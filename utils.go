package gofile

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)


func WalkFileTree(current string, callback func(string, os.FileInfo)) error {
	currentFileInfo, err := os.Stat(current)
	if err != nil {
		return err
	}
	if !currentFileInfo.IsDir() {
		return errors.New("specified path is not directory")
	}
	childFileInfos, err := ioutil.ReadDir(current)
	if err != nil {
		return err
	}
	for _, childFileInfo := range childFileInfos {
		if !childFileInfo.IsDir() {
			callback(current, childFileInfo)
			continue
		}
		err = WalkFileTree(filepath.Join(current, childFileInfo.Name()), callback)
		if err != nil {
			return err
		}
	}
	return nil
}
