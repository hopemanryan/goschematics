package replacementservice

import (
	"fmt"

	argumentsservice "github.com/hopemanryan/goschematics/cmd/arguments-service"
)

func ReplaceFileWithArguments(filePath string, arguments *argumentsservice.FileReplacementArguments) {
	fmt.Printf("%s", filePath)
	fmt.Printf("%s", arguments)

}
