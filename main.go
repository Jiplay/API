package main

import (
	"api/routes"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fmt.Println("Building REST API in Go 1.22")

	mux.HandleFunc("POST /user", routes.CreateUserEndpoint)
	mux.HandleFunc("GET /user", routes.GetUsersEndpoint)
	mux.HandleFunc("DELETE /user", routes.DeleteUserEndpoint)
	mux.HandleFunc("PUT /user", routes.UpdateUserEndpoint)

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
