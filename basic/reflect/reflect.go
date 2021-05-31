package main

import (
	"fmt"
	"reflect"
)

func CustomizeType() {
	type MyInt uint8
	var mi MyInt = 'x'
	v := reflect.ValueOf(mi)
	fmt.Println(reflect.ValueOf(mi))
	fmt.Println(reflect.TypeOf(mi))
	fmt.Println(reflect.ValueOf(mi).Type())
	fmt.Println(reflect.ValueOf(mi).Kind())
	fmt.Println(v.Interface().(MyInt))
}

func ReflectSet() {
	var a uint8 = '0'
	fmt.Println(a)
	v := reflect.ValueOf(&a)
	v.Elem().SetUint('1')
	fmt.Println("a is:", a)
}

func StructReflect() {
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

func main() {
	// var x float64 = 3.14
	// fmt.Println(reflect.TypeOf(x))
	// fmt.Println(reflect.ValueOf(x).String())

	// CustomizeType()

	// ReflectSet()

	StructReflect()
}
