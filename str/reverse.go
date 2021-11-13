package str

import "strings"

// Reverse reverses all the characters in a string and returns the result.
func Reverse(s string) string {
	var strBuilder strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		strBuilder.WriteByte(s[i])
	}
	return strBuilder.String()
}

// ReverseTextfileContents reads the contents of a textfile,
// reverses those contents, and saves them to a destination file.
func ReverseTextfileContents(source string, dest string) error {
	return nil
}
