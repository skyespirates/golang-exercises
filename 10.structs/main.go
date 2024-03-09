package main

import "fmt"

// What is Struct
// Struct  is used to create a collection of members of different data types, into a single variable

type Rectangle struct {
	width 		int
	height		int
}

type Human struct {
	name		string
	age			int
	weight		float32
	height		float32
	isStudent	bool
}

// struct as function argument
func printIdentity(arg Human) {
	fmt.Printf("name\t\t%v\nage\t\t%v\nweight\t\t%v\nheight\t\t%v\nisStudent\t%v\n", arg.name, arg.age, arg.weight, arg.height, arg.isStudent)
}

func main() {
	balok := Rectangle{width: 12, height: 8}
	area := balok.height * balok.width

	var skyes Human = Human{name: "skyes", age: 21, weight: 59.5, height: 173.5, isStudent: false}

	fmt.Println(area)
	printIdentity(skyes)
}