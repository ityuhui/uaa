package main

import (
	"fmt"
	"net/http"
)

func initUserAPIEndpoint() {
	http.HandleFunc("/api/user", userHandler)
	printUserEndpoint()
}

func printUserEndpoint() {
	fmt.Println("")
	fmt.Println("User")
	fmt.Println("POST /api/user Create")
	fmt.Println("PUT /api/user/{id} Update")
	fmt.Println("DELETE /user/{id} Delete")
	fmt.Println("GET /api/user/{id} Get")
	fmt.Println("GET /api/user List")
	fmt.Println("-----------------")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: /api/user")
	switch r.Method {
	case "GET":
		users := StructQueryAllField()
		for _, user := range users {
			fmt.Printf("id: %v, name: %v, password: %v\n", user.Id, user.Name, user.Password)
		}
	}
	AllowCrossDomain(w)
	fmt.Fprintf(w, "Welcome to the /api/user !")

}
