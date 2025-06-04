package main

// Code for assignment 1 -
// https://bjss.learnamp.com/en/learnlists/golang-academy/items/go-assignment
// Create a program that has multiple string variable and
// displays the string on one line. [Strings]

import (
	"fmt"
	"strings"
)

func main() {
	var str1 = "My"
	var str2 = "Name"
	var str3 = "is"
	var str4 = "Martin"

	// use fmt println to print the concatenated strings
	fmt.Println("From println concatenation", str1, str2, str3, str4)

	// alternatively use the strings Builder
	sb := strings.Builder{}
	sb.WriteString("Using strings.Builder..")
	sb.WriteString(str1 + " ")
	sb.WriteString(str2 + " ")
	sb.WriteString(str3 + " ")
	sb.WriteString(str4)

	fmt.Println(sb.String())
}
