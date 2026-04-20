package main

import (
	"fmt"
	"regexp"
)

func fixArticles(s string) string {
	// handles vowels
	re := regexp.MustCompile(`\b([Aa])\s+([aeiouAEIOU]\w*)`)
	s = re.ReplaceAllString(s, "${1}n $2")
	// handles silent "h"
	re2 := regexp.MustCompile(`\b([Aa])\s+(honour|honest|hour|heir)\b`)
	s = re2.ReplaceAllString(s, "$[1]n $2")
	// handles normal "h"
	re3 := regexp.MustCompile(`\b([Aa])n\s+([^aeiouAEIOU]\W*)`)
	s = re3.ReplaceAllString(s, "${1}n $2")

	return s
}
func main() {
	fmt.Println(fixArticles("a amazzing day"))
	fmt.Println(fixArticles("an house there"))
}
