package gofile

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	CannotCopyDirectoryToFileError = errors.New("Cannot copy directory to file")
)

func Copy(src, dst string) error {
	srcFile, err := os.Stat(src)
	if os.IsNotExist(err) {
		return err
	}
	dstFile, err := os.Stat(dst)
	if os.IsNotExist(err) {
		return err
	}
	if srcFile.IsDir() && dstFile.IsDir() {
		return copyRecursive(src, dst)
	} else if !srcFile.IsDir() && dstFile.IsDir() {
		return copyFileToDirectory(src, dst)
	} else if !srcFile.IsDir() && !dstFile.IsDir() {
		return copyFileToFile(src, dst)
	} else {
		return CannotCopyDirectoryToFileError
	}
}

func copyRecursive(src, dst string) error {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(src, file.Name())
		if file.IsDir() {
			dstPath := filepath.Join(dst, file.Name())
			err = os.Mkdir(dstPath, directoryPermission)
			if err != nil {
				return err
			}
			copyRecursive(filePath, dstPath)
			continue
		}
		err = copyFileToDirectory(
			filePath,
			dst,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func copyFileToDirectory(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	base := filepath.Base(src)
	dstFile, err := os.Create(filepath.Join(dst, base))
	if err != nil {
		return err
	}
	defer dstFile.Close()
	bytes, err := ioutil.ReadAll(srcFile)
	if err != nil {
		return err
	}
	_, err = dstFile.Write(bytes)
	return err
}

func copyFileToFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	if !IsExist(dst) {
		file, err := os.Create(dst)
		if err != nil {
			return err
		}
		file.Close()
	}
	dstFile, err := os.Open(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	bytes, err := ioutil.ReadAll(srcFile)
	if err != nil {
		return err
	}
	_, err = dstFile.Write(bytes)
	return err
}
