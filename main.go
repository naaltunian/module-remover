package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var dirName = "node_modules"
	var err error

	path := getPath()

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return filepath.SkipDir
		}
		if info.IsDir() && info.Name() == dirName {
			fmt.Printf("Deleting: %+v \n", info.Name())
			os.RemoveAll(path)
		}
		fmt.Printf("Visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func getPath() string {
	var path string
	// get home directory
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Provide a path to delete all node_modules directories:")
	fmt.Scanln(&path)
	// prevent deleting all node_modules on home path
	if path == homeDir {
		fmt.Println("Cannot use home directory for safety purposes. Enter a different path")
		getPath()
	}
	return path
}
