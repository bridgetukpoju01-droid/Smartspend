package main

import (
	"regexp"
)

func punc(s string) string {
	re := regexp.MustCompile(`\s*([?.£,!:;])`)
	s = re.ReplaceAllString(s, "$1")

	return s
}
package main
import(
	"regexp"

)
func punc(s string) string {
	re := regexp.MustCompile(`\s*([?.,!:;])`)
	s = re.ReplaceAllstrings(s, "$1")
	return s
} package main
import(
	"regexp"
)
func punc(s string) string {
	re := regexp.MustCompile(`\s*([?!.,:;'])`)
	s = re.ReplaceAllStrings(s,"$1")
	return s
}