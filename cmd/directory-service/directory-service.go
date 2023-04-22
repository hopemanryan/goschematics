package directoryservice

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// find all files in directory path that ends with pattern
func FindFilesInDirectory(path string) []string {
	var suffix = "__templ__"
	var files []string

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		if !info.IsDir() && strings.HasSuffix(path, suffix) {
			files = append(files, path)
		}
		return nil
	})

	return files
	// read all files in directory

}

// function that savcs file to directory
func SaveFileToDirectory(path string, file string) {
	// save file to directory
	byte := []byte(file)
	err := os.WriteFile(path, byte, 0644)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

}
