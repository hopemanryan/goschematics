package argumentsservice

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"
)

// struct to hold command line arguments
type FileReplacementArguments struct {
	FileName string `json:"file_name"`
}

func (fr FileReplacementArguments) GetArgumentsMap() map[string]string {
	inInterface := make(map[string]string)
	inrec, _ := json.Marshal(fr)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}

// a function that  get command line variables
func GetArguments() FileReplacementArguments {
	var fileName string
	flag.StringVar(&fileName, "f", "", "file name name to replace (short hand)")
	flag.StringVar(&fileName, "file_name", "", "file name name to replace")

	flag.Parse()

	return FileReplacementArguments{
		FileName: fileName,
	}

}
func (fr FileReplacementArguments) GetArgumentValue(argumentName string) string {
	fmt.Printf("argument: %s \n", argumentName)
	argumentsMap := fr.GetArgumentsMap()
	return argumentsMap[strings.TrimSpace(argumentName)]

}
