package main

import (
	"fmt"
	"log"
	"net/http"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var Employees = make(map[int]Employee)

func main() {

	Employees[1] = Employee{ID: 1, Name: "Raj", Role: "Teacher"}
	Employees[2] = Employee{ID: 2, Name: "Srinadh", Role: "Vice Principal"}

	http.HandleFunc("/greet", Greet)

	log.Fatal(http.ListenAndServe("myuniversity:9000", nil))

}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}

func getAllEmployees(w http.ResponseWriter, r *http.Request) {
	for _, value := range Employees {
		fmt.Fprint(w, value.ID, value.Name, value.Role)
	}
}
