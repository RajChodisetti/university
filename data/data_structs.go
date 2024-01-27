package data

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Major string `json:"major"`
}

var Employees = make(map[int]Employee)
var Students = make(map[int]Student)
