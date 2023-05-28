package student

import (
	"math/rand"
	"strconv"
	"time"
)

type Student struct {
	Id      string `json:"id"`
	Name    string `json:"first_name"`
	Surname string `json:"last_name"`
}

var students []Student

func (student *Student) IsValid() bool {
	if student.Name == "" && student.Surname == "" {
		return false
	}
	return true
}

func IsExist(student Student) bool {
	for _, st := range students {
		if student.Name == st.Name && st.Surname == student.Surname {
			return true
		}
	}
	return false
}
func AddStudent(st *Student) Student {
	//generate id
	rand.Seed(time.Now().UnixNano())
	st.Id = strconv.Itoa(rand.Intn(100))
	students = append(students, *st)
	return *st
}
func UpdateStudent(st Student, id string) {
	s := FindStudent(id)
	DeleteStudent(id)
	st.Id = s.Id
	students = append(students, st)
}

func GetStudents() []Student {
	return students
}

func FindStudent(id string) Student {
	var st Student
	for _, student := range students {
		if student.Id == id {
			st = student
		}
	}
	return st
}
func DeleteStudent(id string) bool {
	for i, student := range students {
		if student.Id == id {
			remove(i)
			return true
		}
	}
	return false
}
func remove(index int) {
	if index < len(students) {
		students = append(students[:index], students[index+1:]...)
	}
}
