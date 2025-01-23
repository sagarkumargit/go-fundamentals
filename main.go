package main

import (
	"fmt"
	"go-fundamentals/helper"
)

func main() {
	helper.FirstFunction()
	const firstname = "Sagar"
	var lastname = "Kumar"
	fmt.Printf("Hello %v %v.\n", firstname, lastname)

	// declare variable with dataTypes

	var fullname string
	var age int
	var email string

	fullname = "Sagar Kumar"
	age = 60
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Printf("He is %v and he is %v years old. Your email is %v.\n", fullname, age, email)

	// CONDITION STATEMENTS
	if age > 50 {
		fmt.Println("Mature")
	} else {
		fmt.Println("Young")
	}
}
