/*
Why not to use some package like Karma https://github.com/reconquest/karma-go for errors and logging these errors?
- It's a third party library
- It's not even popular one (which means time to learn its API)
- Not easy to configure how it format the error stack

Why not to use the well known package https://github.com/pkg/errors
- It's not support structured context by default, so you will have to add support for it yourself.
- It won't work well with the native Wrap errors supported in GO1.13
*/
package main

import (
	"errors"
	"fmt"
	"log"
)

type ContextError struct {
	msg string
	contextMap map[string]interface{}
	cause error
	//callers []interface{}
}

func (contextErr *ContextError) Error() string {
	output := contextErr.msg
	for key, val := range contextErr.contextMap {
		output += fmt.Sprintf("\n%s: %v", key, val)
	}
	err := contextErr.Unwrap()
	output += fmt.Sprintf("\n%s", err)
	return output
}


func (contextErr *ContextError) Unwrap() error {
	return contextErr.cause
}

func createError(msg string, context map[string]interface{}, cause error) error {
	return &ContextError{msg, context, cause}
}

func main() {
	plainErr := errors.New("Plain Error")
	wrappedError := createError("custom wrapped error", nil, plainErr)
	wrappedError = fmt.Errorf("wrappedError %w", wrappedError)
	err := createError("Testing The error with some context", map[string]interface{}{
		"name": "Abdalla",
	}, wrappedError)
	err = createError("Error in Error", nil, err)
	err = createError("third layer", nil, err)
	log.Fatal(err)
}
