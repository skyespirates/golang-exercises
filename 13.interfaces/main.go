package main

import "fmt"

// interface is a way to grouping object based on common method
// object is called implement interface A if and only if that object has all method that interface A has

type person interface {
	getName() string
	getNumber() int
}

// in order to implement interface person, object need implement all method that interface has
// so student need to implement method getName() and getNumber() so the teacher struct
// we can add method to an object using receiver function
type student struct {
	name string
	studentNumber int
}

type teacher struct {
	name string
	teacherNumber int
}

func main() {
	student1 := student{name: "Viola Hunter", studentNumber: 112233}
	teacher1 := teacher{name: "Bessie Greer", teacherNumber: 332211}

	// since types student and teacher satisfy person interface
	// they can be used as argument for print function
	print(student1)
	print(teacher1)

}

func (s student) getName() string {
	return s.name
}
func (s student) getNumber() int {
	return s.studentNumber
}
func (t teacher) getName() string {
	return t.name
}
func (t teacher) getNumber() int {
	return t.teacherNumber
}

// a function that can be called with object that implement person interface
func print(p person) {
	fmt.Println(p.getName())
	fmt.Println(p.getNumber())
}