package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func pathWalkFunc(dir string, fileInfo os.FileInfo, err error) error {
	if fileInfo.IsDir() == false {
		fmt.Println(dir)
	}
	return nil
}

func main() {
	var directory string
	fmt.Println("Enter the directory:")
	fmt.Scanf("%s", &directory)

	err := filepath.Walk(directory, pathWalkFunc)
	if err != nil {
		return
	}
}
