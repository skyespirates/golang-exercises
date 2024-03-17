package main

import "fmt"

// a receiver is a way to embed method to an object
// therefore, all instaces of that type has access to that method
type student struct {
	name string
	age int
}

func main() {
	student1 := student{name: "Elizabeth Kim", age: 20}
	student2 := student{name: "Louisa Johnston", age: 16}
	student1.printIdentity()		// now all student instance has access to method printIdentity
	student2.printIdentity()
}

// create a function that accept a receiver type student
func (s student) printIdentity() {
	fmt.Printf("name is %v, age is %v years old\n", s.name, s.age)
}