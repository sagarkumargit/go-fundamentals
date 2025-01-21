package main

import "fmt"

func main() {
	const firstname = "Sagar"
	var lastname = "Kumar"
	fmt.Printf("Hello %v %v.\n", firstname, lastname)

	// declare variable with dataTypes

	var fullname string
	var age int
	var email string

	fullname = "Sagar Kumar"
	age = 40
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Printf("He is %v and he is %v years old. Your email is %v.\n", fullname, age, email)
}
