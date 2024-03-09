package main

import "fmt"

// What is Maps
// Maps are used to store data in key-value pairs

func main() {
	// create maps
	var myMap1 = map[string]string{"name": "brody", "age": "twelve", "address": "no location"}
	myMap2 := map[string]int{"age": 12, "width": 24, "height": 32, "weight": 64}
	var myMap3 = make(map[string]string)
	myMap4 := make(map[int]int)

	myMap3["name"] = "brody"
	myMap3["age"] = "twentyfour"

	myMap4[1] = 1
	myMap4[2] = 2
	myMap4[3] = 3

	fmt.Println(myMap1["name"])
	fmt.Println(myMap2["width"])
	fmt.Println(myMap3["age"])
	fmt.Println(myMap4[2])
}