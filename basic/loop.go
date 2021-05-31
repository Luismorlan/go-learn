package main

import "log"

func simpleLoop() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}
}

func iterSlice() {
	mySlice := []string{"dog", "cat", "horse"}

	for i, x := range mySlice {
		log.Println(i, x)
	}
}

func iterMap() {
	myMap := make(map[string]string)
	myMap["dog"] = "dog"
	myMap["fish"] = "fish"
	myMap["hat"] = "hat"

	for k, v := range myMap {
		log.Println(k, v)
	}
}

func main() {
	simpleLoop()
	iterMap()
	iterSlice()
}
