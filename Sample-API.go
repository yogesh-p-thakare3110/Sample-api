package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Create a person struct
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Create a address struct
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetpersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	// Create a new router
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Yogesh", Lastname: "Thakare", Address: &Address{City: "Nashik", State: "Maharashtra"}})
	people = append(people, Person{ID: "2", Firstname: "Shikhar", Lastname: "Patil"})
	// Specify Endpoints
	router.HandleFunc("/people", GetpersonEndpoint).Method("GET")
	router.HandleFunc("/people{id}", GetpersonEndpoint).Method("GET")
	router.HandleFunc("/people{id}", CreatPersonEndpoint).Method("POST")
	router.HandleFunc("/people{id}", DeletePersonEndpoint).Method("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}
