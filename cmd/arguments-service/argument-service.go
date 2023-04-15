package argumentsservice

import (
	"encoding/json"
	"flag"
)

// struct to hold command line arguments
type FileReplacementArguments struct {
	FileName string `json:"file_name"`
}

func (fr FileReplacementArguments) GetArgumentsMap() map[string]interface{} {
	var inInterface map[string]interface{}
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
