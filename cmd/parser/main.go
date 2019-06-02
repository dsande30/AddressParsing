/*
	LICENSE GOES HERE

*/

package main

import (
	"fmt"
	"github.com/dsande30/AddressParsing/pkg/parser"
	"github.com/dsande30/AddressParsing/pkg/walker"
	"os"
	"regexp"
)

func main() {
	// Instantiating junk values for testing purposes.
	// This is a regex pointer keyed to "wat"
	re := regexp.MustCompile(".*")
	sampSlice := make(map[string]*regexp.Regexp)
	sampSlice["wat"] = re

	// Create new parser object that contains its own data and expression
	p := parser.Parser{
		Regex: sampSlice,
		Data:  "It was the best of times, it was the worst of times.",
	}
	parser.ParseData(p)

	// Walk directory listed in commandline args
	args := os.Args[1:]
	fmt.Println("Arguments: ", args)
	channel := walker.WalkDir(args[0])
	for msg := range channel {
		fmt.Println(msg)
	}
}
