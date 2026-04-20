package main
import(
	"fmt"
	"strings"
	"strconv"
)
func convertN(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words);i++ {
		if strings.HasPrefix(words[i],"(") && !strings.HasSuffix(words[i],")") {
			a := words[i] + words[i+1]
			a = strings.Trim(a, "()")
			b = strings.Split(a, ",")
			c, _ := strconv.Atoi(b,[1])
			for j := i -c; j < i; j++{
				switch b[0] {
				case "up":
					words[j]= strings.ToUpper(words[j])

				case"low":
					words[j]= strings.ToLower(words[j])

				case"cap":
					words[j]= strings.Title(words[j])
				}

			}
			words = append(words[:i],words[i+1:]...)

		}
	}
	return strings.Join(words," ")
}