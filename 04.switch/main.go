package main

import "fmt"

func main() {
	var num int
	var text string
	fmt.Printf("insert you number:\t")
	fmt.Scan(&num)
	
	// single case
	switch num {
	case 1:
		text = "one"
	case 2:
		text = "two"
	case 3:
		text = "three"
	default:
		text = "not found"
	}
	fmt.Println(text)

	num2 := 7
	// multi case
	switch num2 {
	case 1,3,5:
		fmt.Println("odd number")
	case 2,4,6:
		fmt.Println("even number")
	default:
		fmt.Println("out of bound number")
	}
}