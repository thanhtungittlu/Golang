package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Student struct {
	Id   int
	Name string
	Age  int
}

type Students []Student

func main() {
	fmt.Println("My website is running....")
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	http.HandleFunc("/api/music", MusicApi)
	http.HandleFunc("/api/student", AStudentApi)
	http.HandleFunc("/api/students", ListStudentApi)

	log.Fatal(http.ListenAndServe(":333", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>This is home page<h1>")
}
func AboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>I Love You")
}
func MusicApi(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{
		"name": "At my worst",
		"casi": "Thanh Tung",
	}
	json.NewEncoder(w).Encode(data)
}
func AStudentApi(w http.ResponseWriter, r *http.Request) {
	var data = Student{1, "Tung", 20}
	json.NewEncoder(w).Encode(data)
}
func ListStudentApi(w http.ResponseWriter, r *http.Request) {
	var listStudent = Students{
		Student{1, "Tung", 20},
		Student{1, "Tung2", 22},
		Student{1, "Tung3", 23},
	}
	json.NewEncoder(w).Encode(listStudent)
}
