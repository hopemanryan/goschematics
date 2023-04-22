package main

import (
	"strings"

	argumentsservice "github.com/hopemanryan/goschematics/cmd/arguments-service"
	directoryservice "github.com/hopemanryan/goschematics/cmd/directory-service"
	replacementservice "github.com/hopemanryan/goschematics/cmd/replacement-service"
)

/*
1. save to a specific directory
2. allow config and to run js functions from config
3. allow dynamic file names

*/

func main() {
	arguments := argumentsservice.GetArguments()
	currentDirectoy := directoryservice.GetCurrentDirectory(arguments.ReadDir)
	templateFiles := directoryservice.FindFilesInDirectory(currentDirectoy)

	for _, templateFile := range templateFiles {
		file, _ := replacementservice.ReplaceFileWithArguments(templateFile, &arguments)
		templateFileNameCleanName := strings.Replace(templateFile, "__templ__", "", -1)
		directoryservice.SaveFileToDirectory(templateFileNameCleanName, file)
	}

}
