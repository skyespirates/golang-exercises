package main

import (
	"fmt"
	"os"
)

func main() {
	res, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(res))
}