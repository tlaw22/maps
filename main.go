package main

import "log"

func main() {
	log.Println("Begin...")

	var myString string
	var myInt int

	myString = "Test String"
	myInt = 9

	anotherString := "Another String"
	log.Println(myString, "Is greater than", myInt, "and ", anotherString)

	myMap := make(map[string]string)
	myMap["dog"] = "Sampson"
	myMap["odog"] = "Ricky Lee Bush"

	log.Println(myMap["dog"])
	log.Println(myMap["odog"])
}
