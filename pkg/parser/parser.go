/*
Package parser is responsible for recieving raw data and parsing it.
Essentially it takes a set of regular expressions and searches the given text.
Unsure of how it returns yet. FIXME
*/
package parser

import (
	"fmt"
	"regexp"
)

//Parser responsible for holding regular expressions and parsing given text
type Parser struct {
	// RegEx map[string][]string
	Regex map[string]*regexp.Regexp
	Data  string
}

//ParseData parses its own data based on provided regular expressions
func ParseData(p Parser) {
	fmt.Printf("Given data: %s", p.Data)
}
