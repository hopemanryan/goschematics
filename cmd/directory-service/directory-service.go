package directoryservice

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	argumentsservice "github.com/hopemanryan/goschematics/cmd/arguments-service"
)

func GetCurrentDirectory(basePath string, readDir string) string {

	dir, err := filepath.Abs(basePath + "/" + readDir)

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

func ReadConfigFile(basePath string) map[string]interface{} {
	configFileRaw, err := os.ReadFile(basePath + "/geoschematics.conig.json")
	if err != nil {
		fmt.Println("ERROR:", err)
		return nil
	}

	config := make(map[string]interface{})
	json.Unmarshal(configFileRaw, &config)
	return config

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

func GetConfigByShortHand(shorthand string, entryPath string) map[string]interface{} {
	shortHandValue := ReadConfigFile(entryPath)["shorthands"]
	shortHandConfig := shortHandValue.(map[string]interface{})[shorthand].(map[string]interface{})
	return shortHandConfig

}

func GetWorkingDirectory(arguments argumentsservice.FileReplacementArguments) string {
	if arguments.Shorthand != "" {
		shorthandConfig := GetConfigByShortHand(arguments.Shorthand, arguments.Entry)
		return GetCurrentDirectory(arguments.Entry, shorthandConfig["templatePath"].(string))
	} else {
		return GetCurrentDirectory(arguments.Entry, arguments.ReadDir)
	}

}
