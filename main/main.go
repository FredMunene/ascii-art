package main

import (
	//"ascii"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1:]

	if len(arg) < 1 {
		fmt.Println("Error : offer ")
	}
	
	text := arg[0]
	text2 := strings.Split(text, "\\n")
	text3:= strings.Join(text2, "\n")
 	firstText := text3[:6]
	//secondText := text3[1]
	fmt.Println(text2)
	fmt.Println(text3)
	fmt.Println(string(firstText))
	
	groups := getText()

	line1 := ""
	line2 := ""
	line3 := ""
	line4 := ""
	line5 := ""
	line6 := ""
	line7 := ""
	line8 := ""
	for letter := range firstText {
		key := int(letter)
		if lines, ok := groups[key]; ok {
			if len(lines) > 0 {
				line1 +=lines[0]
				line2 +=lines[1]
				line3 +=lines[2]
				line4 +=lines[3]
				line5 +=lines[4]
				line6 +=lines[5]
				line7 +=lines[6]
				line8 +=lines[7]
			}
		}
	}
		fmt.Println(line1)
		fmt.Println(line2)
		fmt.Println(line3)
		fmt.Println(line4)
		fmt.Println(line5)
		fmt.Println(line6)
		fmt.Println(line7)
		fmt.Println(line8)
}

func getText() map[int][]string {
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error :", err)
	}
	defer file.Close()

	var groupNumber int = 31
	lineCount := 0
	groups := make(map[int][]string)
	// creates a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// skip  empty lines
		if line == "" {
			continue
		}

		if lineCount%8 == 0 {
			groupNumber++
		}
		groups[groupNumber] = append(groups[groupNumber], line)
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
	return groups
}
