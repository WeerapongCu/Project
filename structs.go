package main

import (
	"fmt"
)

type 

func main(){
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	fmt.Println(firstName, lastName, birthdate)
}

func outputUserDetails(firstName , lastName , birthdate string){

	fmt.Println(firstName, lastName, birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}