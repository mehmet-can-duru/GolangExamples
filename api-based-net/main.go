package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func main() {
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/add", addUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users = append(users, user)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User added successfully"))
}
