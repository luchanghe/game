package tool

import (
	"unicode"
)

func FirstToUpper(s string) string {
	if s == "" {
		return ""
	}
	return string(unicode.ToUpper(rune(s[0]))) + s[1:]
}

func FirstToLower(s string) string {
	if s == "" {
		return ""
	}
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
