package replacementservice

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	argumentsservice "github.com/hopemanryan/goschematics/cmd/arguments-service"
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

		return newLine

	} else {
		return line
	}
}
