package main

import (
	argumentsservice "github.com/hopemanryan/goschematics/cmd/arguments-service"
	directoryservice "github.com/hopemanryan/goschematics/cmd/directory-service"
	replacementservice "github.com/hopemanryan/goschematics/cmd/replacement-service"
)

/*
	1. read arguments
	2. get fileName from arguemnts
	3. read file from directory located in by suffix
	4. find places where pattern exits
	5. replace pattern with new value
	6. write new file to directory
*/

func main() {
	currentDirectoy := directoryservice.GetCurrentDirectory()
	templateFiles := directoryservice.FindFilesInDirectory(currentDirectoy)
	arguments := argumentsservice.GetArguments()

	for _, templateFile := range templateFiles {
		replacementservice.ReplaceFileWithArguments(templateFile, &arguments)

	}

}
