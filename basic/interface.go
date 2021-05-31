package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}
	PrintInfo(&dog)
}

func (d *Dog) Says() string {
	return "woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

// Accept Animal.
func PrintInfo(a Animal) {
	log.Println("The animal says", a.Says(), "and has number of legs:", a.NumberOfLegs())
}
