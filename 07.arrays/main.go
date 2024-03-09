package main

import "fmt"

// What is Array?
// Array is a variable that can store multiple values of the same type
// Array also has fixed length

func main() {
	// explicit length
	var arr1 = [3]int{1,2,3}
	arr2 := [4]int{4,5,6,7}

	// implicit length
	var arr3 = [...]string{"hello", "world"}
	arr4 := [...]bool{true, false}


	fmt.Println(arr1, arr2, arr3, arr4)

	// get array length
	fmt.Println(len(arr1), len(arr2), len(arr3), len(arr4))

}