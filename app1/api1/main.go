package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func greet(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Method is Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	courses := []string{"java", "C#.Net", "SQL", "React"}
	json.NewEncoder(res).Encode(courses)
}

func getCourses(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Metod is Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	courses := []string{"java", "C#.Net", "SQL", "React"}
	json.NewEncoder(res).Encode(courses)
}

type Course struct {
	Code string `json:code`
	Name string `json:code`
}

func addCourseHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Method is Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	defer req.Body.Close()

	var course Course

	err := json.NewDecoder(req.Body).Decode(&course)
	if err != nil {
		http.Error(res, "Invalid Data", http.StatusBadRequest)
	}

	fmt.Printf("Received: %+v\n", course)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(map[string]string{
		"message":    "Record Saved Successfully",
		"courseName": course.Name,
	})
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/courses", getCourses)
	http.HandleFunc("/addcourses", addCourseHandler)

	fmt.Println("Server is up and running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
