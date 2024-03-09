package main

import "fmt"

// What is Slice?
// Slice is similiar to an array but more powerful
// whereas length of slice could grow or shrink

func main() {
	// create slice
	students := []string{"badu", "budi", "heru"}
	
	// declare empty slice
	mySlice := []int{}

	fmt.Println(mySlice)

	fmt.Printf("%v\t%v\n", len(students), cap(students))
	
	// create slice from an array
	myArr := [...]int{1,2,3,4,5,6}
	myslice := myArr[1:4]
	fmt.Println(myslice)

	// create slice using make() function
	sliceWithMake := make([]int, 5, 10)
	fmt.Printf("%v\t%v\n%v\n", len(sliceWithMake), cap(sliceWithMake), sliceWithMake)

	// modify slice using append() function
	sliceWithMake = append(sliceWithMake, 12,13,14)
	fmt.Printf("%v\t%v\n%v\n", len(sliceWithMake), cap(sliceWithMake), sliceWithMake)
	for i:=15; i<20; i++ {
		sliceWithMake = append(sliceWithMake, i)
	}
	fmt.Printf("%v\t%v\n%v\n", len(sliceWithMake), cap(sliceWithMake), sliceWithMake)
}