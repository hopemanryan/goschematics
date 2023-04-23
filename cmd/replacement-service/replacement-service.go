package replacementservice

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	argumentsservice "github.com/hopemanryan/goschematics/cmd/arguments-service"
	v8 "rogchap.com/v8go"
)

func ReplaceFileWithArguments(filePath string, arguments *argumentsservice.FileReplacementArguments) (newFileData string, error bool) {
	lines := []string{}
	arguments.GetArgumentsMap()
	var newFile = ""
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return "", true
	} else {

		for _, line := range strings.Split(string(data), "\n") {
			newLine := ReplaceLine(line, arguments)
			lines = append(lines, newLine)
		}

		newFile = strings.Join(lines, "\n")

	}

	return newFile, false

}

// write a function that checks if a string has a regex of anything between <% and =%> and replace it with the value of the argument
func ReplaceLine(line string, arguments *argumentsservice.FileReplacementArguments) string {
	replaceableRegex := regexp.MustCompile(`(?m)<%.*=%>`)

	// if the line has a regex match
	if replaceableRegex.MatchString(line) {

		// regex of group between <% and =%>
		regex := regexp.MustCompile(`(?m)<%(.*)=%>`)
		// get the group
		group := regex.FindStringSubmatch(line)[1]
		// get the value of the argument
		argValue := arguments.GetArgumentValue(group)

		newLine := replaceableRegex.ReplaceAllString(line, argValue)
		RunJS()

		return newLine

	} else {
		return line
	}
}

func RunJS(function string) {

	configRaw, _ := os.ReadFile("goschematics.js")
	script := string(configRaw)

	iso := v8.NewIsolate()
	ctx := v8.NewContext(iso)

	_, err := ctx.RunScript(script, "goschematics.js")
	if err != nil {
		e := err.(*v8.JSError)    // JavaScript errors will be returned as the JSError struct
		fmt.Println(e.Message)    // the message of the exception thrown
		fmt.Println(e.Location)   // the filename, line number and the column where the error occured
		fmt.Println(e.StackTrace) // the full stack trace of the error, if available

		fmt.Printf("javascript error: %v", e)        // will format the standard error message
		fmt.Printf("javascript stack trace: %+v", e) // will format the full error stack trace
	}

	val, _ := ctx.RunScript(function, "goschematics.js")
	fmt.Println(val.String())

}
