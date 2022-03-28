package main

import "fmt"

//Insert variables declarations here based on activity

func tellMeTypes() {
	//insert code here to print out types of variables
	text := "The following is the account information"
	firstName := "Luke"
	age := 20
	credit := 123.55
	subscribe := true
	fmt.Printf("%T %T %T %T %T", text, firstName, age, credit, subscribe)
}

func main() {
	tellMeTypes()
}
