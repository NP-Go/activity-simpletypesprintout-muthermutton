package main

import "fmt"

//Insert variables declarations here based on activity

func tellMeTypes() {
	//insert code here to print out types of variables
	text := "The following is the account information"
	firstName := "Luke"
	lastName := "Skywalker"
	age := 20
	weight := 73.0
	height := 1.72
	credit := 123.55
	accName := "admin"
	accPassword := "password"
	subscribe := true
	fmt.Printf("%T %T %T %T %T %T %T %T %T %T", text, firstName, lastName, age, weight, height, credit, accName, accPassword, subscribe)
}

func main() {
	tellMeTypes()
}
