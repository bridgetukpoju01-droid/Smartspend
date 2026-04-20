package main

import (
	"fmt"
	"regexp"
)

func punc(s string) string {
	re := regexp.MustCompile(`\s*([?.£,!:;])`)
	s = re.ReplaceAllString(s, "$1")

	return s
}
func main() {
	fmt.Println(punc(" we are happy ! so happy ?"))
}
