package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("You are an adult.")
	}else if age >= 13 {
		fmt.Println("You are a teenager.")
	}else {
		fmt.Println("You are a minor.")
	}


	var role = "admin"
	var hasPermission = true


	if role == "admin" || hasPermission {
		fmt.Println("Access granted.")
	}


	if age := 15 ; age >= 18 {
		fmt.Println("You are an adult.")
	}	else {
		fmt.Println("You are not an adult.")
	}	

	// No ternary operator in Go, use if-else instead
}