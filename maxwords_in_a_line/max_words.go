package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "the quick brown fox jumps over the lazy dog"

	// Tokenise the string with space as delimiter
	toklist := strings.Split(str, " ")

	// Number of characters in a line
	lineWidth := 20

	var tempstr string
	lines := make([]string, 0)

	// Parse and form lines that fit within line-width
	for _, tok := range toklist {
		if len(tempstr)+len(tok) < lineWidth {
			if len(tempstr) == 0 {
				tempstr = tok
			} else {
				tempstr = tempstr + " " + tok
			}
		} else {
			lines = append(lines, tempstr)
			tempstr = tok
		}
	}
	lines = append(lines, tempstr)

	for _, l := range lines {
		fmt.Printf("%s(%d)\n", l, len(l))
	}
}
