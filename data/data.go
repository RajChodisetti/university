package data

var e = 3
var s = 3

func Datahouse() {
	Employees[1] = Employee{ID: 1, Name: "Raj", Role: "Teacher"}
	Employees[2] = Employee{ID: 2, Name: "Praveena", Role: "Principal"}

	Students[1] = Student{ID: 101, Name: "Srinadh", Major: "CS"}
	Students[2] = Student{ID: 203, Name: "Madhu", Major: "Yes"}
}

func AddEmployee(Emp Employee) {
	Employees[e] = Emp
	e++
}

func AddStudent(Stu Student) {
	Students[s] = Stu
	s++
}
