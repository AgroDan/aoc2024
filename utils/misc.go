package utils

import "strings"

// This is just a bunch of functions that are simple but can't really be categorized
// in a larger part

func ReverseString(s string) string {
	// This will reverse a string
	ret := make([]rune, len(s))
	for i, r := range s {
		ret[len(s)-i-1] = r
	}
	return string(ret)
}

// zero fill, like what python does
func ZFill(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return strings.Repeat("0", width-len(s)) + s
}
