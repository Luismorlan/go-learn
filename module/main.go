package main

import (
	"log"

	"example.com/myprog/helpers"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Type Name"
	log.Println(myVar)
}
