package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RajChodisetti/university/data"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	for _, value := range data.Employees {
		fmt.Fprintf(w, "ID: %d, Name: %s, Role: %s\n", value.ID, value.Name, value.Role)
	}
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	empID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	employee, ok := data.Employees[empID]
	if !ok {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	employeeJSON, err := json.Marshal(employee)
	if err != nil {
		http.Error(w, "Error encoding employee to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(employeeJSON)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	empID, err := strconv.Atoi(idStr)
	if err == nil {
		delete(data.Employees, empID)
	}
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	// Get values from query parameters
	idStr := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	role := r.URL.Query().Get("role")

	// Convert ID to integer
	empID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID: %s", idStr), http.StatusBadRequest)
		return
	}

	// Check if the employee already exists
	_, exists := data.Employees[empID]
	if exists {
		http.Error(w, "Employee with the same ID already exists", http.StatusBadRequest)
		return
	}

	// Create a new Employee struct with the provided data
	newEmployee := data.Employee{
		ID:   empID,
		Name: name,
		Role: role,
	}

	// Add the new employee to the map
	data.Employees[empID] = newEmployee

	// Respond with the added employee details
	employeeJSON, err := json.Marshal(newEmployee)
	if err != nil {
		http.Error(w, "Error encoding employee to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(employeeJSON)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	// Get values from query parameters
	idStr := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	role := r.URL.Query().Get("role")

	// Convert ID to integer
	empID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID: %s", idStr), http.StatusBadRequest)
		return
	}
	employee := data.Employees[empID]
	if name != "" {
		employee.Name = name
	}
	if role != "" {
		employee.Role = role
	}
	data.Employees[empID] = employee
	employeeJSON, err := json.Marshal(employee)
	if err != nil {
		http.Error(w, "Error encoding to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(employeeJSON)

}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	for _, value := range data.Students {
		fmt.Fprintf(w, "ID: %d, Name: %s, Major: %s\n", value.ID, value.Name, value.Major)
	}
}
