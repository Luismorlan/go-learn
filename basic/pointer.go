package main

import "log"

func main() {
	myString := "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("After my function call, myString is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("pointer s is set to", s)
	*s = "Red"
}
