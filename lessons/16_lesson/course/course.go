package course

import (
	"math/rand"
	"myapiexample/student"
	"strconv"
	"time"
)

type Course struct {
	Id       string            `json:"id"`
	Name     string            `json:"name"`
	Students []student.Student `json:"students,omitempty"`
}
type CourseStudent struct {
	CourseId  string `json:"courseid"`
	StudentId string `json:"studentid"`
}
type defCourse struct {
	Id       string
	Name     string
	Students []string
}

var courses []defCourse

func FindCourse(id string) Course {
	var course Course
	for _, c := range courses {
		if c.Id == id {
			course = changeCoure(c)
		}
	}
	return course
}

func (course *Course) IsValid() bool {
	return course.Name != ""
}

func IsExist(course Course) bool {
	for _, c := range courses {
		if course.Name == c.Name {
			return true
		}
	}
	return false
}
func AddCourse(course *Course) Course {
	//generate id
	rand.Seed(time.Now().UnixNano())
	var students = []string{}
	for _, st := range course.Students {
		students = append(students, st.Id)
	}
	c := defCourse{
		Id:       strconv.Itoa(rand.Intn(100)),
		Name:     course.Name,
		Students: students,
	}
	courses = append(courses, c)
	return changeCoure(c)
}
func UpdateCourse(cr Course, id string) {
	c := FindCourse(id)
	var students = []string{}
	for _, st := range c.Students {
		students = append(students, st.Id)
	}
	DeleteCourse(id)
	newCourse := defCourse{
		Id:       c.Id,
		Name:     cr.Name,
		Students: students,
	}
	courses = append(courses, newCourse)
}

func AddStudent(data CourseStudent) bool {
	st := student.FindStudent(data.StudentId)
	if st.Id == "" {
		return false
	}

	//generate id
	for index, c := range courses {
		if c.Id == data.CourseId {
			for _, id := range c.Students {
				if data.StudentId == id {
					return false
				}
			}
			remove(index)
			course := defCourse{
				Id:       c.Id,
				Name:     c.Name,
				Students: append(c.Students, data.StudentId),
			}
			courses = append(courses, course)
			return true
		}
	}
	return false

}

func GetCourses() []Course {
	return changeCoureses()
}
func DeleteCourse(id string) bool {
	for i, course := range courses {
		if course.Id == id {
			remove(i)
			return true
		}
	}
	return false
}

// default function
func changeCoureses() []Course {
	var ls = []Course{}
	for _, c := range courses {
		ls = append(ls, changeCoure(c))
	}
	return ls
}

func changeCoure(c defCourse) Course {
	stList := []student.Student{}

	for _, stId := range c.Students {
		st := student.FindStudent(stId)
		stList = append(stList, st)
	}

	return Course{
		Id:       c.Id,
		Name:     c.Name,
		Students: stList,
	}
}

func remove(index int) {
	if index < len(courses) {
		courses = append(courses[:index], courses[index+1:]...)
	}
}
