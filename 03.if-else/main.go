package main

import "fmt"

func main() {
	isActive := false
	if !isActive {
		fmt.Println("yey active")
	} else {
		fmt.Println("unactive :(")
	}
	day := 7
	if day==1 {
		fmt.Println("sunday")
	} else if day==2 {
		fmt.Println("monday")
	} else if day==3 {
		fmt.Println("tuesday")
	} else if day==4 {
		fmt.Println("wednesday")
	} else if day==5 {
		fmt.Println("thursday")
	} else if day==6 {
		fmt.Println("friday")
	} else if day==7 {
		fmt.Println("saturday")
	} else {
		fmt.Println("Please insert day between 1-7")
	}
}