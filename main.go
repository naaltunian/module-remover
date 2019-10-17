package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	var dirName = "node_modules"
	var path string
	var moduleDirPath string
	var err error

	flag.StringVar(&moduleDirPath, "path", "none", "File path to be used")
	flag.Parse()

	if moduleDirPath == "none" {
		path = getPath()
	} else {
		path = moduleDirPath
	}

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

// prompts the user for path if not given
func getPath() string {

	var path string
	// get home directory
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Provide a path to delete all node_modules directories:")
	fmt.Scanln(&path)

	validPath := checkPath(path)

	// prevent deleting all node_modules on home path
	if path == homeDir {
		fmt.Println("Cannot use home directory for safety purposes. Enter a different path:")
		path = getPath()
	} else if !validPath {
		fmt.Println("Provide a valid path:")
		path = getPath()
	}
	return path
}

// check if path exists
func checkPath(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
