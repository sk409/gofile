package gofile

import "os"

func makeFile(path, text string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(text))
	return err
}
