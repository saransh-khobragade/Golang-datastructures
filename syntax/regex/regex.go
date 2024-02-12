package main

import (
	"fmt"
	"regexp"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func main() {
	str := "Test@%String#321gosamples.dev ـا ą ٦"
	fmt.Println(clearString(str))
}
