package util

import (
	"fmt"
	"io/ioutil"
)

func GetFilesAndDirectories(path string) (string, []string, []string, error) {
	filenames := []string{}
	directories := []string{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "", filenames, directories, fmt.Errorf("%+v", err)
	}

	for _, f := range files {
		if f.IsDir() {
			directories = append(directories, f.Name())
		} else {
			filenames = append(filenames, f.Name())
		}
	}

	return path, directories, filenames, nil
}
