package main

import "log"

func main() {
	// var whatToSay string
	var i int

	whatToSay, _ := saySomething("hello world")

	i = 5
	log.Println(whatToSay)
	log.Println(i)
}

func saySomething(s string) (string, string) {
	return s, "Say " + s
}
