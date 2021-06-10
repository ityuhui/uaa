package main

import (
	"encoding/json"
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
	AllowCrossDomain(w)
	switch r.Method {
	case "GET":
		users := UserQueryAllField()
		for _, user := range users {
			fmt.Printf("id: %v, name: %v, password: %v\n", user.Id, user.Name, user.Password)
		}
		bytes, _ := json.Marshal(users)
		fmt.Fprint(w, string(bytes))
	}
}
