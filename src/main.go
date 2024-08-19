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
	// Routes
	s.HandleFunc("/getAllOrders", getAllOrders).Methods("GET")
	s.HandleFunc("/getOrder/{id}", getOrder).Methods("GET")

	s.HandleFunc("/createProfile", createProfile).Methods("POST")
	s.HandleFunc("/getUserProfile", getUserProfile).Methods("POST")
	s.HandleFunc("/updateProfile", updateProfile).Methods("PUT")
	s.HandleFunc("/deleteProfile/{id}", deleteProfile).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}
