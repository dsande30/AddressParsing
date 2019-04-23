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
	re := regexp.MustCompile(".*")
	sampSlice := make(map[string]*regexp.Regexp)
	sampSlice["wat"] = re
	fmt.Println(sampSlice)
	p := parser.Parser{sampSlice, "at"}
	parser.ParseData(p)

	args := os.Args[1:]
	fmt.Println(args)
	walker.WalkDir(".")
}
