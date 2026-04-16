package main

import (
	"fmt"
	"strings"
)

func convertcase(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		switch strings.ToLower(words[i]) {
		case "(up)":
			for j := 0; j < i; j++ {
				words[j] = strings.ToUpper(words[j])
			}
			words = append(words[:i], words[i+1:]...)
			i--
		case "(low)":
			for j := 0; j < i; j++ {
				words[j] = strings.ToLower(words[j])
			}
			words = append(words[:i], words[i+1:]...)
			i--
		case "(cap)":
			for j := 0; j < i; j++ {
				if len(words[j]) > 0 {
					words[j] = strings.ToUpper(string(words[j][0])) + strings.ToLower(words[j][1:])
				}
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(convertcase("Welcome guys (up)"))
	fmt.Println(convertcase("hello , WORLD (low)"))
	fmt.Println(convertcase("we are happy now (cap)"))
}
