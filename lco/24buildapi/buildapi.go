package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// model for course
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// helper file
func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - LCO")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{
		CourseId:    "9399453b-4dff-41c0-8a77-e06ca2f630a7",
		CourseName:  "GO Lang",
		CoursePrice: 1999,
		Author: &Author{
			"Nitin Kumar",
			"go.dev",
		},
	})

	// routing
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getCourse).Methods("GET")
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")

	// listen on a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controller file
// serve home route
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to APIs by LCO</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a Courses")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, c := range courses {
		if c.CourseId == params["id"] {
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a Courses")
	w.Header().Set("Content-Type", "application/json")

	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil || course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON: " + err.Error())
		return
	}

	// generate unique id, string
	// append course into courses

	course.CourseId = uuid.New().String()
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update a Courses")
	w.Header().Set("Content-Type", "application/json")

	// grad id from request
	params := mux.Vars(r)
	id := params["id"]

	// loop and find the course from courses array
	// remove it from the array
	// then add the course received from the request body
	for index, course := range courses {
		if course.CourseId == id {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = id
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the provided ID")
	return
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete a Courses")
	w.Header().Set("Content-Type", "application/json")

	// grad id from request
	params := mux.Vars(r)
	id := params["id"]

	// loop and find the course from courses array
	// remove it from the array
	for index, course := range courses {
		if course.CourseId == id {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the provided ID")
	return
}
