package argumentsservice

import (
	"encoding/json"
	"flag"
	"strings"
)

// struct to hold command line arguments
type FileReplacementArguments struct {
	FileName  string `json:"file_name"`
	ReadDir   string `json:"read_dir"`
	OutDir    string `json:"out_dir"`
	Shorthand string `json:"shorthand"`
	Entry     string `json:"entry"`
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
	var readDir string
	var outDir string
	var shorthand string
	var entryPoint string
	flag.StringVar(&fileName, "file_name", "", "file name name to replace")
	flag.StringVar(&readDir, "read_dir", "", "directory to read files from")
	flag.StringVar(&outDir, "out_dir", "", "directory to write files to")
	flag.StringVar(&shorthand, "shorthand", "", "Shorthand for predefind in configf file")
	flag.StringVar(&shorthand, "s", "", "Shorthand for predefind in configf file")
	flag.StringVar(&entryPoint, "_entrypoint", "", "directory to write files to")

	flag.Parse()

	return FileReplacementArguments{
		FileName:  fileName,
		ReadDir:   readDir,
		OutDir:    outDir,
		Entry:     entryPoint,
		Shorthand: shorthand,
	}

}
func (fr FileReplacementArguments) GetArgumentValue(argumentName string) string {
	argumentsMap := fr.GetArgumentsMap()
	return argumentsMap[strings.TrimSpace(argumentName)]

}
