package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func mapDemo() {
	// This stores anything in the map:
	// make(map[string]interface{})
	myMap := make(map[string]User)

	myMap["me"] = User{
		FirstName: "Alex",
		LastName:  "Bella",
	}

	log.Println(myMap["me"])
}

func sliceDemo() {
	var myString []string
	myString = append(myString, "Warren")
	myString = append(myString, "Mary")

	sort.Strings(myString)
	log.Println(myString)

	numbers := []int{4, 3, 2, 1}
	log.Println(numbers)

	// Partially print out slice.
	log.Println(numbers[0:2])
}

func main() {
	mapDemo()
	sliceDemo()
}
