package main

import "fmt"

func main() {

	var i, j string = "Hello", "World"
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	var (
		name string = "skyes"
		age int = 32
		isManager bool = false
		growth float32 = 1.12
	)

	fmt.Println(name, age, isManager, growth)

	// %T -> print data type
	// %v -> print value in default format
	// %#v -> print in GO-syntax format
	var a int = 12
	fmt.Printf("%T, %v,  %#v\n", a, a, a)
	var b string = "kiwkiw cukurukuk"
	fmt.Printf("%T, %v, %#v\n", b, b, b)
	var c bool = true
	fmt.Printf("%T, %v, %#v\n", c, c, c)
	var d float32 = 3.14
	fmt.Printf("%T, %v, %#v\n", d, d, d)

}