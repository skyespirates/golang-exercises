package main

// constants outside function
const pi float32 = 3.14

func main() {
	// declare varible
	var name string = "skyes"
	var age int = 32

	// declare multiple variables at once
	var str1, str2, str3 string = "hoho", "hehe", "wkwk"
	var isActive, isCompleted bool = true, false

	var (
		a, b, c int = 1, 2, 3
		d, e, f float32 = 4.4, 5.5, 6.6
		g, h string = "ggg", "hhh"
	)

	// constants
	const (
		constA int = 12
		constB = "hola"
		constC bool = false
	)

}