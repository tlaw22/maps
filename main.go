package main

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

	myMap := make(map[string]string)
	myMap["dog"] = "Sampson"
	myMap["odog"] = "Ricky Lee Bush"

	log.Println(myMap["dog"])
	log.Println(myMap["odog"])
	log.Println("================================")
	myMap2 := make(map[string]int)
	myMap2["First"] = 6
	myMap2["Second"] = 3
	log.Println(myMap2["First"], "Plus", myMap2["Second"], "Equals: ", myMap2["First"]+myMap2["Second"])

	type User struct {
		FirstName string
		LastName  string
	}
	myMap3 := make(map[string]User)
	me := User{
		FirstName: "Tim",
		LastName:  "Law",
	}
	myMap3["me"] = me
	log.Println(myMap3["me"].FirstName)
	log.Println(myMap3["me"].LastName)

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

}
