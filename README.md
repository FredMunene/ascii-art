ASCII-ART PROGRAM
=============

This is part of the projects I am currently doing in learning Go.

The program receives a string as an argument and outputting the string in a graphic representation using ASCII.

`go run ."Hello" | cat -e`

# Requirements for the Program

+ The program handles an input with numbers, letters, spaces, special characters and '\n' .
+ The program is written in Go.
+ The code respects the [good practices.](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/good-practices/README.md)
+ There are test files for [unit testing](https://go.dev/doc/tutorial/add-a-test).
+ [Use Standard Go](https://golang.org/pkg/) packages.
+ **Banner files** with a specific graphical template represenation using ASCII are given. The files are pre-formatted so no need to change them. 
    * [shadow](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt)
    * [standard](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt)
    * [thinkertoy](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt)


## Banner Format
================

+ Each character has a height of 8 lines.
+ Characters are seperated by a new line '\n'.
