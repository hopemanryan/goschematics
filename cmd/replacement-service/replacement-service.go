package replacementservice

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	argumentsservice "github.com/hopemanryan/goschematics/cmd/arguments-service"
	jsservice "github.com/hopemanryan/goschematics/cmd/js-service"
)

func ReplaceFileWithArguments(filePath string, arguments *argumentsservice.FileReplacementArguments, jsContext *jsservice.JSContext) (newFileData string, error bool) {
	lines := []string{}
	arguments.GetArgumentsMap()
	var newFile = ""
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return "", true
	} else {

		jsContext.SetJSPreFunctions()
		for _, line := range strings.Split(string(data), "\n") {
			newLine := ReplaceLine(line, arguments, jsContext)
			lines = append(lines, newLine)
		}

		newFile = strings.Join(lines, "\n")

	}

	return newFile, false

}

// write a function that checks if a string has a regex of anything between <% and =%> and replace it with the value of the argument
func ReplaceLine(line string, arguments *argumentsservice.FileReplacementArguments, jsContext *jsservice.JSContext) string {
	replaceableRegex := regexp.MustCompile(`(?m)<%.*=%>`)

	// if the line has a regex match
	if replaceableRegex.MatchString(line) {

		// regex of group between <% and =%>
		regex := regexp.MustCompile(`(?m)<%(.*)=%>`)
		// get the group
		group := regex.FindStringSubmatch(line)[1]

		if strings.Contains(group, "(") && strings.Contains(group, ")") {

			insiteParenthesis := regexp.MustCompile(`(?m)\((.*)\)`)
			innerValue := insiteParenthesis.FindStringSubmatch(group)[1]
			argValue := arguments.GetArgumentValue(innerValue)
			stringedValue := fmt.Sprintf("'%s'", argValue)
			jsFunction := strings.Replace(group, "("+innerValue+")", "("+stringedValue+")", -1)
			funcResult := jsContext.RunScript(jsFunction)
			newLineFunctionResultLine := replaceableRegex.ReplaceAllString(line, funcResult)
			return newLineFunctionResultLine

		} else {
			fmt.Println("no function")

			argValue := arguments.GetArgumentValue(group)

			newLine := replaceableRegex.ReplaceAllString(line, argValue)

			return newLine
		}
		// get the value of the argument

	} else {
		return line
	}
}
