package jsservice

import (
	"fmt"
	"os"

	"rogchap.com/v8go"
)

type JSContext struct {
	ctx *v8go.Context
}

func NewJSFunction() *JSContext {
	iso := v8go.NewIsolate()

	return &JSContext{
		ctx: v8go.NewContext(iso),
	}
}

func (js *JSContext) SetJSPreFunctions() {
	configRaw, _ := os.ReadFile("goschematics.js")
	script := string(configRaw)
	_, err := js.ctx.RunScript(script, "goschematics.js")
	if err != nil {
		e := err.(*v8go.JSError)  // JavaScript errors will be returned as the JSError struct
		fmt.Println(e.Message)    // the message of the exception thrown
		fmt.Println(e.Location)   // the filename, line number and the column where the error occured
		fmt.Println(e.StackTrace) // the full stack trace of the error, if available

		fmt.Printf("javascript error: %v", e) // will format the standard error message
		fmt.Printf("javascript stack trace: %+v", e)
		panic(e.Message) // will format the full error stack trace
	}

}

func (js *JSContext) RunScript(functionScript string) string {
	val, scriptErr := js.ctx.RunScript(functionScript, "goschematics.js")
	if scriptErr != nil {
		e := scriptErr.(*v8go.JSError) // JavaScript errors will be returned as the JSError struct
		fmt.Println(e.Message)         // the message of the exception thrown
		fmt.Println(e.Location)        // the filename, line number and the column where the error occured
		fmt.Println(e.StackTrace)
		panic(e.Message) // will format the full error stack trace

	}

	return val.String()

}
