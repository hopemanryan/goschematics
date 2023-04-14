package argumentsservice

import (
	"flag"
)

// struct to hold command line arguments
type FileReplacementArguments struct {
	file_name string
}

// a function that  get command line variables
func GetArguments() FileReplacementArguments {
	var fileName string
	flag.StringVar(&fileName, "f", "", "file name name to replace (short hand)")
	flag.StringVar(&fileName, "file_name", "", "file name name to replace")

	flag.Parse()

	return FileReplacementArguments{fileName}

}
