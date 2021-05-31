package main

import "log"

func simpleIfDemo() {
	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println(isTrue)
	} else {
		log.Println(isTrue)
	}
}

func aBitMoreFlavor() {
	cat := "cat"
	if cat == "cat" {
		log.Println("Cat is cat")
	}
}

func numberDemo() {
	myNum := 100
	isTrue := true
	if myNum > 99 && isTrue {
		log.Println("myNum is greater than 99 and isTrue == True")
	}
}

func switchDemo() {
	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("is Cat")
	case "dog":
		log.Println("is Dog")
	default:
		log.Println("is Nothing at all")
	}
}

func main() {
	simpleIfDemo()
	aBitMoreFlavor()
	numberDemo()
	switchDemo()
}
