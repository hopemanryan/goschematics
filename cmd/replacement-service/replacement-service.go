package replacementservice

import (
	"fmt"
	"os"
	"strings"

	argumentsservice "github.com/hopemanryan/goschematics/cmd/arguments-service"
)

func ReplaceFileWithArguments(filePath string, arguments *argumentsservice.FileReplacementArguments) (newFileData string, error bool) {
	arguments.GetArgumentsMap()
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return "", true
	} else {

		for i, line := range strings.Split(string(data), "\n") {
			fmt.Println(i, line)
		}

	}

	return "", false

}

func FindInContent(text string) {

}
