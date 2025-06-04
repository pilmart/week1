package main

// Code for assignment 2 -
// https://bjss.learnamp.com/en/learnlists/golang-academy/items/go-assignment
// Create a program that lets the user input a first name, middle name and last name.
// Display the person's full name on one line. [Keyboard input]

import (
	"fmt"
	"strings"
)

func main() {

	var firstName string
	var middleName string
	var lastName string

	// use fmt scan to print the concatenated strings
	fmt.Println("Please enter first name")
	// retrieve the value from the user
	fmt.Scan(&firstName)

	fmt.Println("Please enter your middle name")
	// retrieve the value from the user
	fmt.Scan(&middleName)

	fmt.Println("Please enter last name")
	// retrieve the value from the user
	fmt.Scan(&lastName)

	// alternatively use the strings Builder
	sb := strings.Builder{}
	sb.WriteString("Your name is : ")
	sb.WriteString(firstName + " ")
	sb.WriteString(middleName + " ")
	sb.WriteString(lastName + " ")

	fmt.Println(sb.String())
}
