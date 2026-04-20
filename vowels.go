package main

import (
	"fmt"
	"regexp"
)

func fixArticle(s string) string {
	// vowels
	re1 := regexp.MustCompile(`\b([Aa])\s+([aeiouAEIOU]\w*)`)
	s = re1.ReplaceAllString(s, "${1}n $2")

	// silent 'h' words
	re2 := regexp.MustCompile(`\b([Aa])\s+(hour|honest|honor|heir)\b`)
	s = re2.ReplaceAllString(s, "${1}n $2")

	return s
}

func main() {
	fmt.Println(fixArticle("A amazing rock!"))
	fmt.Println(fixArticle("A hour later"))
	fmt.Println(fixArticle("A house nearby"))
}
