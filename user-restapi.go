package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func initUserAPIEndpoint() {
	rtr.HandleFunc("/api/user/{id:[0-9]+}", userHandler)
	http.Handle("/", rtr)
	printUserEndpoint()
}

func printUserEndpoint() {
	fmt.Print(`User
POST /api/user Create
PUT /api/user/{id} Update
PUT /api/user/{id} Update
PUT /api/user/{id} Update
DELETE /user/{id} Delete
GET /api/user/{id} Get
GET /api/user List
-----------------
`)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: /api/user")
	AllowCrossDomain(w)
	switch r.Method {
	case "POST":
		createUser()
	case "DELETE":
		deleteUser()
	case "GET":
		users := UserQueryAllField()
		for _, user := range users {
			fmt.Printf("id: %v, name: %v, password: %v\n", user.Id, user.Name, user.Password)
		}
		bytes, _ := json.Marshal(users)
		fmt.Fprint(w, string(bytes))
	}
}
