package main

import "fmt"

func main() {
	// normal loop
	for i:=0; i<3; i++ {
		fmt.Printf("%v ", i)
	}

	// loop increment by 10
	fmt.Print("\n")
	for i:=0; i<=100; i+=10 {
		fmt.Printf("%v ", i)
	}
	fmt.Print("\n")

	// continue
	for i:=0; i<=100; i+=10 {
		if(i>10 && i<90) {
			continue
		}
		fmt.Printf("%v ", i)
	}
	fmt.Print("\n")

	// break
	for i:=0; i<=100; i+=10 {
		if(i==50) {
			break
		}
		fmt.Printf("%v ", i)
	}
	fmt.Print("\n")

	// nested loop
	for i:=0; i<5; i++ {
		for j:=0; j<=i; j++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}

	// looping through slice
	// slice is kind of array but with flexible length, array has fixed length btw
	var intArray = []int{1,2,3}
	for idx, val := range intArray {
		fmt.Printf("%v\t%v\n", idx, val)
	}
	strArray := []string{"aaa", "bbb", "ccc", "ddd"}
	for _, value := range strArray {
		fmt.Printf("%v\t%v\n", value, value)
	}
}