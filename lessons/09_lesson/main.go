package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		//create 1
		var s Student
		//default values
		fmt.Println(s)
		//create 2
		t := new(Teacher)
		//default values
		fmt.Println("pointer", t)
		fmt.Println(*t)
		//studen speak
		s.Speak()
		//teacher speak
		t.Speak()
		//Student change age
		fmt.Println("before student age", s.age)
		s.ChangeAge(5)
		fmt.Println("after student age", s.age)

		//crete 3
		t2 := Teacher{"Saidmurod", "Turdiyev", 24, "developer", 1}
		fmt.Println("Teacher 2 ", t2)

		//crete 4
		s2 := Student{name: "Saidmurod", surname: "Turdiyev", age: 24, course: 0, index: 1}
		fmt.Println("Student 2 ", s2)

		sub := Subject{Student: s, Teacher: *t, name: "Math", index: 0}

		fmt.Println("Subject ", sub)

	*/

	var circle Circle = Circle{radius: 5.3}
	rectangle := Rectangle{x: 4, y: 6}

	showPeremetre(circle)
	showPeremetre(rectangle)
}

/*
type Student struct {
	name    string
	surname string
	age     int
	course  int
	index   int
}
type Teacher struct {
	name    string
	surname string
	age     int
	subject string
	index   int
}
type Subject struct {
	Student
	Teacher
	name  string
	index int
}
type Talk interface {
	Speak()
	ChangeAge(age int)
}

func (s Student) Speak() {
	fmt.Println("Student", s.name, "is speaking")
}

func (t Teacher) Speak() {
	fmt.Println("Teacher", t.name, "is speaking")
}

//	func (s Student) ChangeAge(age int) {
//		s.age = age
//		fmt.Println("Student age", s.age)
//		fmt.Println("Changed age  ", age)
//		fmt.Println("but base one inside main function did not change you must use pointer")
//	}
func (s *Student) ChangeAge(age int) {
	s.age = age
	fmt.Println("Student age", s.age)
	fmt.Println("Changed age  ", age)
}
*/

type Circle struct {
	radius float64
}
type Rectangle struct {
	x float64
	y float64
}

func (c Circle) peremetre() float64 {
	return math.Pi * 2 * c.radius
}

func (r Rectangle) peremetre() float64 {
	return 2 * (r.x + r.y)
}

func showPeremetre(shape Shape) {
	fmt.Println("this shape peremetre is", shape.peremetre())
}

type Shape interface {
	peremetre() float64
}
