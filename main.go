package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	Db = db()
	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path

	// Apply the CORS middleware
	http.Handle("/", CORS(route))

	// Routes
	s.HandleFunc("/getAllOrders", getAllOrders).Methods("GET")
	s.HandleFunc("/getOrder/{id}", getOrder).Methods("GET")
	s.HandleFunc("/addOrder", addOrder).Methods("POST")
	s.HandleFunc("/chargeBackOrder", chargeBackOrder).Methods("POST")
	s.HandleFunc("/validateOrders", isValidOrders).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}
