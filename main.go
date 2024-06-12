package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "This is a test. This testis only a test."
	word := "test"

	// Compile the regular expression
	re := regexp.MustCompile(word)

	// Find all matches
	matches := re.FindAllStringIndex(text, -1)

	if len(matches) == 0 {
		fmt.Println("No matches found.")
	} else {
		lastIndex := 0
		fmt.Printf("Found %d matches:\n", len(matches))
		
		for _, match := range matches {
			start, end := match[0], match[1]
			if start > lastIndex {
				fmt.Printf("%s", text[lastIndex:start])
			}
			fmt.Printf("\x1b[31m%s\x1b[0m", text[start:end])

			lastIndex = end
		}
		if lastIndex < len(text) {
			fmt.Printf("%s", text[lastIndex:])
		}
		fmt.Println()
	}
}
