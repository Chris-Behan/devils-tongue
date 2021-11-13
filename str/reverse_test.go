package str

import (
	"testing"
)

func TestReverse(t *testing.T) {
	originalStr := "hello world!"
	reversedStr := "!dlrow olleh"
	result := Reverse(originalStr)
	if result != reversedStr {
		t.Fatalf("Expected reverse string to be %v but it was %v", result, reversedStr)
	}
}
