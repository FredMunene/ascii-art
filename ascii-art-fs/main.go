package main

import (
	"ascii-art-fs/functions"
	"fmt"
	"os"
)

func main() {
	// flag if no fileName is provided as argument
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER] ")
		fmt.Println("Ex: go run . \"Hello\" standard")
		return
	}
	// gets input string to interpret
	input := args[0]

	// gets the name of the banner file to usego
	bannerFile := args[1]

	// checks Banner name is correctly handled
	bannerFileName := functions.CheckBannerName(bannerFile)

	// checks input string is  valid and handled
	str := functions.ReadInput(input)

	// checks strings against ascii figures and assigns to string
	result := functions.PrintAsciiArt(str, bannerFileName)

	fmt.Print(result)
}
