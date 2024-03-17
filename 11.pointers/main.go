package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	skyes := person{name: "skyes", age: 22}
	skyes.setAge(36)			// without pointer the change wont reflect
	skyes.setName("hulk")		// with pointer to modify the original object
	fmt.Println(skyes)
}

func (p person) setAge(newAge int) {
	p.age = newAge
}

func (p *person) setName(newName string) {
	p.name = newName
}