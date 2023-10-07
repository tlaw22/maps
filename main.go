package main

// This lesson teaches maps and slices

import (
	"log"
	"sort"
)

func main() {
	log.Println("Begin...")

	var myString string
	var myInt int

	myString = "Test String"
	myInt = 9

	anotherString := "Another String"
	log.Println(myString, "Is greater than", myInt, "and ", anotherString)

	// Declare a map
	myMap := make(map[string]string)
	myMap["dog"] = "Sampson"
	myMap["odog"] = "Ricky Lee Bush"

	log.Println(myMap["dog"])
	log.Println(myMap["odog"])
	log.Println("================================")
	// Do math with a map
	myMap2 := make(map[string]int)
	myMap2["First"] = 6
	myMap2["Second"] = 3
	log.Println(myMap2["First"], "Plus", myMap2["Second"], "Equals: ", myMap2["First"]+myMap2["Second"])
	// declare a struct that holds 2 strings
	// Structs seem to function a lot like simple classes

	type User struct {
		FirstName string
		LastName  string
	}
	myMap3 := make(map[string]User)

	// Assign the struct to be stored in a ne variable that can be used later
	me := User{
		FirstName: "Tim",
		LastName:  "Law",
	}
	myMap3["me"] = me

	// Print out the struct portions

	log.Println(myMap3["me"].FirstName)
	log.Println(myMap3["me"].LastName)

	// Slices work the same way as arrays

	var mySlice []string
	mySlice = append(mySlice, "Florida")
	mySlice = append(mySlice, "Georgia")
	mySlice = append(mySlice, "Texas")
	mySlice = append(mySlice, "Utah")
	mySlice = append(mySlice, "New Jersey")

	log.Println(mySlice)

	var mySlice2 []int

	mySlice2 = append(mySlice2, 5)
	mySlice2 = append(mySlice2, 6)
	mySlice2 = append(mySlice2, 7)
	mySlice2 = append(mySlice2, 0)
	mySlice2 = append(mySlice2, 9)

	log.Println(mySlice2)
	sort.Ints(mySlice2)
	log.Println(mySlice2)
	// declare a slive in one line
	numbers := []int{5, 33, 77, 12, 74, 90}
	log.Println(numbers)
	log.Println(numbers[0:2])
}
