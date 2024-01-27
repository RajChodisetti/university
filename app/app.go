package app

import (
	"log"
	"net/http"
)

func app() {
	data.datahouse()
	http.HandleFunc("/greet", Greet)
	http.HandleFunc("/getEmployees", getAllEmployees)

	log.Fatal(http.ListenAndServe("myuniversity:9000", nil))

}
