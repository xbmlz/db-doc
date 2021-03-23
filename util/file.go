package util

import (
	"io/ioutil"
	"log"
	"os"
)

// CreateDir
func CreateDir(dirPath string) {
	if !IsExist(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		os.Chmod(dirPath, os.ModePerm)
		log.Println(err.Error())
		os.Exit(1)
	} else {
		// TODO tips
	}
}

// IsExist
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// WriteToFile write file
func WriteToFile(path, content string) {
	//syscall.Umask(0000)
	if err := ioutil.WriteFile(path, []byte(content), 0777); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
