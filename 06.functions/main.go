package main

import "fmt"

func stringPrinter(arg string) {
	fmt.Println(arg)
}

func trianglePattern(height int) {
	for i:=0; i<height; i++ {
		for j:=0; j<=i; j++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
}

func main() {
	stringPrinter("hello world")
	trianglePattern(4)
	fmt.Println(add(2,3))
	fmt.Println(mul(3,4))
}

// before or after main function doesn't matter
func add(num1 int, num2 int) int {
	return num1 + num2
}

func mul(num1 int, num2 int) int {
	return num1 * num2
}