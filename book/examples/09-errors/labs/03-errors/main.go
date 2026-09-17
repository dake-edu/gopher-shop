package main

import (
	"errors"
	"fmt"
)

type titleError struct{ reason string }

func (problem titleError) Error() string {
	return problem.reason
}
func main() {
	var absent error
	fmt.Println(absent == nil)
	var problem error = titleError{reason: "empty title"}
	fmt.Println(problem.Error())
	cause := errors.New("missing file")
	wrapped := fmt.Errorf("read title: %w", cause)
	fmt.Println(wrapped)
	fmt.Println(errors.Is(wrapped, cause))
	fmt.Println(errors.Is(wrapped, errors.New("missing file")))
}
