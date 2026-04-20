package main

import (
	"strings"
)

func vowels(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words)-1; i++ {
		vowel := "aeiouhAEIOUH"
		next := words[i+1]
		switch words[i] {
		case "A":
			if strings.ContainsAny(string(next[0]), vowel) {
				words[i] = "An"
			}
		case "a":
			if strings.ContainsAny(string(next[0]), vowel) {
				words[i] = "an"

			}
		}
	}
	return strings.Join(words, " ")
}
package main
import (
	"strings"
)
func vowels(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words)-1 i++{
		vowel := "aeiouhAEIOUH"
		next := words[i+1]
		switch := words[i] {
		case"A":
			if strings.ContainsAny(string(next[0]), vowel) {
				words[i] "An"
			case"a":
				if strings.ContainAny(String(next[0]), vowel) {

				}
			}
		}
	}
} package main
import(
	"strings"
)
func vowels(s string) string {
	words := strings.Fields(s)
	for i := 0 -c; i < len(words)-1; i++ {
		vowel := "aeiouhAEIOUH"
		next := words[i+1]
		switch := words[i] {
		case "A":
			if strings.ContainAny(strings(next[0]),vowels) {
				words[i] "An"

		case "a":
			if srings.ContainAny(strings(next[0]), vowels) {
				words[i] "an"
			}
			

			}
			
		}

	}
	return strings.Join(words," ")
	} 
