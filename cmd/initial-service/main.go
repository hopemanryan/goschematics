package main

import (
	"strings"

	argumentsservice "github.com/hopemanryan/goschematics/cmd/arguments-service"
	directoryservice "github.com/hopemanryan/goschematics/cmd/directory-service"
	jsservice "github.com/hopemanryan/goschematics/cmd/js-service"
	replacementservice "github.com/hopemanryan/goschematics/cmd/replacement-service"
)

/*
1.implement save to directory
2. support generic config file
*/

func main() {
	arguments := argumentsservice.GetArguments()
	currentDirectoy := directoryservice.GetWorkingDirectory(arguments)
	templateFiles := directoryservice.FindFilesInDirectory(currentDirectoy)
	jsContext := jsservice.NewJSFunction()
	for _, templateFile := range templateFiles {
		file, _ := replacementservice.ReplaceFileWithArguments(templateFile, &arguments, jsContext)
		cleanFileName := replacementservice.ReplaceLine(templateFile, &arguments, jsContext)
		templateFileNameCleanName := strings.Replace(cleanFileName, "__templ__", "", -1)
		directoryservice.SaveFileToDirectory(templateFileNameCleanName, file)
	}
}
