package str

import (
	"bytes"
	"os"
	"testing"
)

func TestReverseString(t *testing.T) {
	originalStr := "hello world!"
	reversedStr := "!dlrow olleh"
	result := ReverseString(originalStr)
	if result != reversedStr {
		t.Fatalf("Expected reverse string to be %v but it was %v", reversedStr, result)
	}
}

func TestReverseBytes(t *testing.T) {
	originalBytes := []byte{'a', 'b', 'c'}
	reversedBytes := []byte{'c', 'b', 'a'}
	result := ReverseBytes(originalBytes)
	if !bytes.Equal(result, reversedBytes) {
		t.Fatalf("Expected reverse bytes to be %v but it was %v", reversedBytes, result)
	}
}

func TestReverseTextFileContents(t *testing.T) {
	contents := "hello \n world"
	expectedContents := "dlrow \n olleh"
	sourceFile := "test.txt"
	os.WriteFile(sourceFile, []byte(contents), 0666)

	resultsFile := "results.txt"
	ReverseTextfileContents(sourceFile, resultsFile)

	resultBytes, err := os.ReadFile(resultsFile)
	if err != nil {
		t.Fatalf("Error reading results file: %v", err)
	}
	result := string(resultBytes)
	if result != expectedContents {
		t.Fatalf("Expected result to be %v but it was %v", expectedContents, result)
	}

	// Clean up files created by test
	os.Remove(sourceFile)
	os.Remove(resultsFile)
}
