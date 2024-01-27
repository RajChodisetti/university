package app

import (
	"log"
	"net/http"

	"github.com/RajChodisetti/university/data"
	"github.com/RajChodisetti/university/handler"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()
	data.Datahouse()

	router.HandleFunc("/greet", handler.Greet)

	router.HandleFunc("/getemployees", handler.GetAllEmployees)
	router.HandleFunc("/employees/get", handler.GetEmployees)
	router.HandleFunc("/employees/update", handler.UpdateEmployee)
	router.HandleFunc("/employees/delete", handler.DeleteEmployee)
	router.HandleFunc("/employees/add", handler.AddEmployee)

	router.HandleFunc("/getstudents", handler.GetAllStudents)

	log.Fatal(http.ListenAndServe("myuniversity:9000", nil))

}
