package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Person struct {
	ID                string   `json:"id,omitempty"`
	Firstname         string   `"json:firstname,omitempty"`
	Lastname          string   `"json:lastname,omitempty"`
	PreviousCTC       string   `json:"previousctc,omitempty"`
	ExpectedCTC       string   `json:"expectedctc,omitempty"`
	Willingtorelocate string   `json:"willingtorelocate,omitempty"`
	Skills            string   `json:"skills,omitempty"`
	Overallexp        string   `json:"overallexp,omitempty"`
	Relevantexp       string   `json:"relevantexp,omitempty"`
	Address           *Address `"json:address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitemty"`
	State string `json:"state,omitempty"`
}

var jobseekers []Person

func getPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobseekers)

}
func getPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for _, item := range jobseekers {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func createPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var people Person
	_ = json.NewDecoder(r.Body).Decode(&people)
	people.ID = strconv.Itoa(rand.Intn(1000000))
	jobseekers = append(jobseekers, people)
	json.NewEncoder(w).Encode(people)
}
func updatePersonEdPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range jobseekers {
		if item.ID == params["id"] {
			jobseekers = append(jobseekers[:index], jobseekers[index+1:]...)
			var people Person
			_ = json.NewDecoder(r.Body).Decode(&people)
			people.ID = params["id"]
			jobseekers = append(jobseekers, people)
			json.NewEncoder(w).Encode(people)
			return
		}
	}
	json.NewEncoder(w).Encode(jobseekers)

}
func deletePersonEdpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range jobseekers {
		if item.ID == params["id"] {
			jobseekers = append(jobseekers[:index], jobseekers[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(jobseekers)

}

func main() {
	// init router
	r := mux.NewRouter()

	jobseekers = append(jobseekers, Person{ID: "001", Firstname: "Gnana Naga Surya", Lastname: "Parvathi", PreviousCTC: "300000 PA", ExpectedCTC: "600000 PA", Willingtorelocate: " yes", Skills: "GO,SQL, RESTAPI, Github, POSTMAN", Overallexp: " 2.5 years", Relevantexp: "0", Address: &Address{City: "Kakinada", State: "Adhra Pradesh"}})
	jobseekers = append(jobseekers, Person{ID: "002", Firstname: "Karthik", Lastname: "Krishna", PreviousCTC: "350000 PA", ExpectedCTC: "650000 PA", Willingtorelocate: " yes", Skills: "GO, SQL, RESTAPI, Github, POSTMAN, NODE.JS, ReactNative", Overallexp: " 4 years", Relevantexp: "1", Address: &Address{City: "Hyderabad", State: "Telangana"}})

	r.HandleFunc("/api/jobseekers", getPersonEndpoint).Methods("GET")
	r.HandleFunc("/api/jobseekers/{id}", getPeopleEndpoint).Methods("GET")
	r.HandleFunc("/api/jobseekers", createPersonEndpoint).Methods("POST")
	r.HandleFunc("/api/jobseekers/{id}", updatePersonEdPoint).Methods("PUT")
	r.HandleFunc("/api/jobseekers/{id}", deletePersonEdpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
