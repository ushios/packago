package packago

import "fmt"

func Concat(a string, b string) string {
	var text string
	text = fmt.Sprintf("%s%s", a, b)
	return text
}
