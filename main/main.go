package main

import (
	//"ascii"
	"bufio"
	"fmt"
	"os"
)

func main() {
	// ascii.Printer()
	// ascii.PrintGrimReaper("Hello")
	// opening a file
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error :", err)
	}
	defer file.Close()
	// read the data into a slice of bytes
	/*
		data := make([]byte, 1000) // initialises a slice of type byte, and length 1000.
		count, err := file.Read((data))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("read %d bytes: %q\n", count, data[:count])
	*/

	scanner := bufio.NewScanner(file)

	for i := 0; i<20 && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
}
