package iteration

import "strings"

var repeatCount = 5

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

// Additional function for flexibility
func RepeatCustom(character string, count int) string {
	var repeated strings.Builder
	for i := 0; i < count; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
