package main

import "os"

func main() {
	str := "hey hey not bad"
	os.WriteFile("textFile.txt", []byte(str), 0666)
}