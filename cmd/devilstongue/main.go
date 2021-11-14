package main

import (
	"flag"
	"fmt"

	"github.com/Chris-Behan/devils-tongue/str"
)

func main() {

	flag.Usage = func() {
		fmt.Println("Reverse a string or use the input and output flags to reverse the contents of a file.\nEx. devilstongue 'string to reverse' or devilstongue -input 'input.txt' -output 'reversed.txt'")
		flag.PrintDefaults()
	}
	var inputFilename string
	var outputFilename string
	flag.StringVar(&inputFilename, "input", "", "Name of the input file whose contents you want to reverse. If this flag is used, you must also use the output flag.")
	flag.StringVar(&outputFilename, "output", "", "Name of the output file you want to write the reversed contents to. If this flag is used, you must also use the input flag.")
	flag.Parse()

	if inputFilename == "" && outputFilename == "" {
		positionalArgs := flag.Args()
		if len(positionalArgs) != 1 {
			fmt.Print("Expected a single string argument. Ex: 'string to reverse'")
			return
		}

		inputString := positionalArgs[0]
		reversedString := str.ReverseString(inputString)
		fmt.Print(reversedString)
	} else {
		if inputFilename == "" {
			fmt.Print("Please specify the name of the input file. Ex: '-input input.txt'")
			return
		}
		if outputFilename == "" {
			fmt.Print("Please specify the name of the output file. Ex: '-output output.txt'")
			return
		}
		err := str.ReverseTextfileContents(inputFilename, outputFilename)
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Printf("Reversed the contents of %v and saved them in %v", inputFilename, outputFilename)
	}
}
