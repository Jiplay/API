package routes

import (
	"api/jwt"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

import "api/user"

func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Printf("Create user")
	var newUser user.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token := newUser.Create()
	_, _ = fmt.Fprintf(w, token)
}

func UpdateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Printf("Update user")
	var newUser user.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newUser.Update()
	_, _ = fmt.Fprintf(w, "ok")
}

func DeleteUserEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Printf("Delete user")

	clientToken := r.Header.Get("Token")
	_, err := jwt.ValidateToken(clientToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var newUser user.User
	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newUser.Delete()
	_, _ = fmt.Fprintf(w, "ok")
}

func GetUsersEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get users")
	user.Get()
}
