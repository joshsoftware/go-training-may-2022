package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var users = []User{
	{ID: 1, FirstName: "Shubham", LastName: "Vyas"},
	{ID: 2, FirstName: "John", LastName: "Doe"},
	{ID: 3, FirstName: "Nancy", LastName: "Walters"},
	{ID: 4, FirstName: "Charlie", LastName: "Smith"},
	{ID: 5, FirstName: "Chris", LastName: "Rock"},
	{ID: 6, FirstName: "Cris", LastName: "Black"},
}

func getUserDetails(w http.ResponseWriter, r *http.Request) {
	// id will be in path
	// We can get id by removing every thing before id
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/user/"))
	// validating id
	if err != nil || id > len(users)-1 || id == 0 {
		fmt.Println("Invalid id passed")
		http.Error(w, "Invalid id", 400)
		return
	}
	var user User
	for _, v := range users {
		if v.ID == id {
			user = v
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// use can access this api at
// http://localhost:8080/users
func getUsers(w http.ResponseWriter, r *http.Request) {
	// Send json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/user/", getUserDetails)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
