package main

import (
	"fmt"
	"strconv"
	"strings"
)

func baseconvert(s string) string {
	word := strings.Fields(s)
	for i := 1; i < len(word); i++ {
		if word[i] == "(hex)" {
			dec, err := strconv.ParseInt(word[i-1], 16, 64)
			if err != nil {
				fmt.Println("Not a valid number")
				continue
			}
			word[i-1] = strconv.FormatInt(dec, 10)
			word = append(word[:i], word[i+1:]...)
		} else if word[i] == "(bin)" {
			dec, err := strconv.ParseInt(word[i-1], 2, 64)
			if err != nil {
				fmt.Println("not a valid number")
				continue
			}
			word[i-1] = strconv.FormatInt(dec, 10)
			word = append(word[:i], word[i+1:]...)
		}
	}
	return strings.Join(word, " ")
}

func main() {
	fmt.Println(baseconvert("we have 1E (hex) bags for sale"))
	fmt.Println(baseconvert("1010 (bin) dresses"))
}
