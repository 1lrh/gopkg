package fsx

import (
	"os"
	"path/filepath"
)

func CreateDirIfNotExist(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func CheckDirExist(path string) bool {
	f, err := os.Stat(path)
	return err == nil && f.IsDir()
}

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
