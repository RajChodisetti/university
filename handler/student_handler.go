package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RajChodisetti/university/data"
)

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	for _, value := range data.Students {
		fmt.Fprintf(w, "ID: %d, Name: %s, Major: %s\n", value.ID, value.Name, value.Major)
	}
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	stuId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	Student, ok := data.Students[stuId]
	if !ok {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	StudentJSON, err := json.Marshal(Student)
	if err != nil {
		http.Error(w, "Error encoding Student to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(StudentJSON)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	stuId, err := strconv.Atoi(idStr)
	if err == nil {
		delete(data.Students, stuId)
	}
}

func AddStudent(w http.ResponseWriter, r *http.Request) {
	// Get values from query parameters
	idStr := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	major := r.URL.Query().Get("major")

	// Convert ID to integer
	stuId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID: %s", idStr), http.StatusBadRequest)
		return
	}

	// Check if the Student already exists
	_, exists := data.Students[stuId]
	if exists {
		http.Error(w, "Student with the same ID already exists", http.StatusBadRequest)
		return
	}

	// Create a new Student struct with the provided data
	newStudent := data.Student{
		ID:    stuId,
		Name:  name,
		Major: major,
	}

	// Add the new Student to the map
	data.Students[stuId] = newStudent

	// Respond with the added Student details
	StudentJSON, err := json.Marshal(newStudent)
	if err != nil {
		http.Error(w, "Error encoding Student to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(StudentJSON)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	// Get values from query parameters
	idStr := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	major := r.URL.Query().Get("major")

	// Convert ID to integer
	stuId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID: %s", idStr), http.StatusBadRequest)
		return
	}
	Student := data.Students[stuId]
	if name != "" {
		Student.Name = name
	}
	if major != "" {
		Student.Major = major
	}
	data.Students[stuId] = Student
	StudentJSON, err := json.Marshal(Student)
	if err != nil {
		http.Error(w, "Error encoding to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(StudentJSON)

}
