package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertbase(s string) string {
	words := strings.Fields(s)
	for i := 1; i < len(words); i++ {
		switch strings.ToLower(words[i]) {
		case "(hex)":
			if dec, err := strconv.ParseInt(words[i-1], 16, 64); err == nil {
				words[i-1] = strconv.FormatInt(dec, 10)
				words = append(words[:i], words[i+1:]...)

			}
		case "(bin)":
			if dec, err := strconv.ParseInt(words[i-1], 2, 64); err == nil {
				words[i-1] = strconv.FormatInt(dec, 10)
				words = append(words[:i], words[i+1:]...)

			}

		}
	}
	return strings.Join(words, " ")

}
func main() {
	fmt.Println(convertbase("hello 1E (hex) people"))
	fmt.Println(convertbase("welcome 1010 (bin) connections"))
}
