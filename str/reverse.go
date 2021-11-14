package str

import (
	"os"
	"strings"
)

// ReverseString reverses all the characters in a string and returns the result.
func ReverseString(s string) string {
	var strBuilder strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		strBuilder.WriteByte(s[i])
	}
	return strBuilder.String()
}

// ReverseBytes reverses all bytes in a byte slice and returns the result.
func ReverseBytes(b []byte) []byte {
	for left, right := 0, len(b)-1; left < right; left, right = left+1, right-1 {
		b[left], b[right] = b[right], b[left]
	}
	return b
}

// ReverseTextfileContents reads the contents of a textfile,
// reverses those contents, and saves them to a destination file.
func ReverseTextfileContents(source string, dest string) error {
	fileContents, readErr := os.ReadFile(source)
	if readErr != nil {
		return readErr
	}

	reversedContents := ReverseBytes(fileContents)

	writeErr := os.WriteFile(dest, reversedContents, 0666)
	if writeErr != nil {
		return writeErr
	}

	return nil
}
