package main

import (
	"encoding/json"
	"fmt"
	"log"
	"myapiexample/course"
	"myapiexample/student"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	runSerever()
}

func runSerever() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", addCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PATCH")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	r.HandleFunc("/courseaddstudent", addStudentToCourse).Methods("POST")

	r.HandleFunc("/students", getAllStudents).Methods("GET")
	r.HandleFunc("/student/{id}", getStudent).Methods("GET")
	r.HandleFunc("/student", addStudent).Methods("POST")
	r.HandleFunc("/student/{id}", updateStudent).Methods("PATCH")
	r.HandleFunc("/student/{id}", deleteStudent).Methods("DELETE")

	//
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Run Server")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welocme to golang series on my example</h1>"))
}

// get all
func getAllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all students")
	w.Header().Set("Content-Type", "application/json")
	students := student.GetStudents()
	json.NewEncoder(w).Encode(students)
	return
}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all course")
	w.Header().Set("Content-Type", "application/json")
	courses := course.GetCourses()
	json.NewEncoder(w).Encode(courses)
	return
}

// get one
func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	c := course.FindCourse(params["id"])
	if c.Id == "" {
		json.NewEncoder(w).Encode("Course is not find")
		return
	}
	json.NewEncoder(w).Encode(c)
	return
}
func getStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Student")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	s := student.FindStudent(params["id"])
	if s.Id == "" {
		json.NewEncoder(w).Encode("Student is not find")
		return
	}
	json.NewEncoder(w).Encode(s)
	return
}

// add
func addCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	var newCourse course.Course
	_ = json.NewDecoder(r.Body).Decode(&newCourse)
	if !newCourse.IsValid() {
		json.NewEncoder(w).Encode("You send wrong data \ndata:" + newCourse.Name)
		return
	}
	if !course.IsExist(newCourse) {
		c := course.AddCourse(&newCourse)
		json.NewEncoder(w).Encode(c)
	} else {
		json.NewEncoder(w).Encode("This data is exist")

	}
	return

}
func addStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Student")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	var newStudent student.Student
	_ = json.NewDecoder(r.Body).Decode(&newStudent)
	if !newStudent.IsValid() {
		json.NewEncoder(w).Encode("You send wrong data")
		return
	}
	if !student.IsExist(newStudent) {
		neSt := student.AddStudent(&newStudent)
		json.NewEncoder(w).Encode(neSt)
	} else {
		json.NewEncoder(w).Encode("This data is exist")

	}
	return
}
func addStudentToCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Student to course")
	w.Header().Set("Content-Type", "application/json")
	var crSt course.CourseStudent
	_ = json.NewDecoder(r.Body).Decode(&crSt)

	isAdded := course.AddStudent(crSt)
	if isAdded {
		c := course.FindCourse(crSt.CourseId)
		json.NewEncoder(w).Encode(c)
	} else {
		json.NewEncoder(w).Encode("Data is not added" + crSt.CourseId)
	}
	return
}

// update
func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	c := course.FindCourse(params["id"])
	if c.Name != "" && c.Id != "" {
		var newCourse course.Course
		_ = json.NewDecoder(r.Body).Decode(&newCourse)
		course.UpdateCourse(newCourse, params["id"])
		json.NewEncoder(w).Encode(newCourse)
	} else {
		json.NewEncoder(w).Encode("No course found")
	}
	return
}
func updateStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Student")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	s := student.FindStudent(params["id"])
	if !s.IsValid() {
		var newStudent student.Student
		_ = json.NewDecoder(r.Body).Decode(&newStudent)
		student.UpdateStudent(newStudent, params["id"])
		json.NewEncoder(w).Encode(newStudent)
	} else {
		json.NewEncoder(w).Encode("No student found")
	}
	return
}

// delete
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	isDelete := course.DeleteCourse(params["id"])

	if isDelete {
		json.NewEncoder(w).Encode("Delete")
	} else {
		json.NewEncoder(w).Encode("No course found")
	}
	return
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete student")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	isDelete := student.DeleteStudent(params["id"])

	if isDelete {
		json.NewEncoder(w).Encode("Delete")
	} else {
		json.NewEncoder(w).Encode("No Student found")
	}
	return
}
