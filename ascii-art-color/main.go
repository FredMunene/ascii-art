package main

import (
	"ascii-art-fs/functions"
	"fmt"
	"os"
	// "regexp"
	"strings"
)

func main() {
	// flag if no fileName is provided as argument
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")

	}

	// obtaining user input
	color := os.Args[1]
	// lettersToColor := os.Args[2]
	words := os.Args[3]
	bannerFileName := "standard.txt"

	// spliting input color string to extract the color name
	colorText := strings.Split(color,"=")
	colorName := colorText[1]

	// obtainng RGB values for color using 'PrintColored' function
	r,g,b := PrintColored(colorName)

	// concatenates everything into a string
	foregroundColor := fmt.Sprintf("\x1b[38;2;%d;%d;%dm",r,g,b)
	reset_color := "\x1b[0m"

	// Regular expression pattern to match the word
	// pattern := lettersToColor

	// compile the regular expression
	// re := regexp.MustCompile(pattern)

	// // Check if the word exists
	// if re.MatchString(words){
	// 	fmt.Printf("%s %v %s\n",foregroundColor,lettersToColor,reset_color)
	// } else {
	// 	fmt.Printf("%s %v %s\n",foregroundColor,words,reset_color)
	// }
	words = functions.ReadInput(words) + "\n"
	fmt.Println(words)

	result := functions.PrintAsciiArt(words, bannerFileName)

	fmt.Printf("%s %v %s",foregroundColor,result,reset_color)
}
