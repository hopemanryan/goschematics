package replacementservice

import (
	"fmt"
	"os"

	argumentsservice "github.com/hopemanryan/goschematics/cmd/arguments-service"
)

func ReplaceFileWithArguments(filePath string, arguments *argumentsservice.FileReplacementArguments) (newFileData string, error bool) {
	argJson := arguments.GetArgumentsMap()
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return "", true
	} else {
		fmt.Printf("%s", argJson)
		fmt.Println(string(data))

	}

	return "", false

}
