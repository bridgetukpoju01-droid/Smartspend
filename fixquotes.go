package main

import (
	"fmt"
	"strings"
)

func fixquotes(text string) string {
	// handle single quotes
	text = strings.ReplaceAll(text, "' ", "'")
	text = strings.ReplaceAll(text, " '", "'")
	// handles double quotes
	text = strings.ReplaceAll(text, "\"", "\"")
	text = strings.ReplaceAll(text, "\" ", "\"")

	return text
}
func main() {
	fmt.Println(fixquotes(" ' we are happy ' "))
	fmt.Println(fixquotes("\" where are you\""))
}
