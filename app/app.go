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
	http.Handle("/", router)

	router.HandleFunc("/greet", handler.Greet)

	router.HandleFunc("/getemployees", handler.GetAllEmployees)
	router.HandleFunc("/employees/get", handler.GetEmployee).Methods("GET")
	router.HandleFunc("/employees/update", handler.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employees/delete", handler.DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/employees/add", handler.AddEmployee).Methods("POST")

	router.HandleFunc("/getstudents", handler.GetAllStudents)

	log.Fatal(http.ListenAndServe("myuniversity:9000", nil))

}
