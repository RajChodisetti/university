package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RajChodisetti/university/data"
	"github.com/gorilla/mux"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	for _, value := range data.Employees {
		fmt.Fprintf(w, "ID: %d, Name: %s, Role: %s\n", value.ID, value.Name, value.Role)
	}
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empID, err := strconv.Atoi(vars["id"])
	if err == nil {
		fmt.Fprint(w, data.Employees[empID])
	}
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empID, err := strconv.Atoi(vars["id"])
	if err == nil {
		delete(data.Employees, empID)
	}
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empID, err := strconv.Atoi(vars["id"])
	if err == nil {
		var newEmployee data.Employee
		err = json.NewDecoder(r.Body).Decode(&newEmployee)
		if err == nil {
			data.Employees[empID] = newEmployee
		}
	}
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empID, err := strconv.Atoi(vars["id"])
	if err == nil {
		var updatedEmployee data.Employee
		err = json.NewDecoder(r.Body).Decode(&updatedEmployee)
		if err == nil {
			data.Employees[empID] = updatedEmployee
		}
	}
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	for _, value := range data.Students {
		fmt.Fprintf(w, "ID: %d, Name: %s, Major: %s\n", value.ID, value.Name, value.Major)
	}
}
