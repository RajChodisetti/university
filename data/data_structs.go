package data

type Employee struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	Role string `json:"role" xml:"role"`
}

type Student struct {
	ID    int    `json:"id" xml:"id"`
	Name  string `json:"name" xml:"name"`
	Major string `json:"major" xml:"major"`
}

var Employees = make(map[int]Employee)
var Students = make(map[int]Student)
