package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func initUserAPIEndpoint(r *gin.Engine) {
	r.GET("/api/user", listUserHandler)
	r.GET("/api/user/:id", getUserHandler)
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

func listUserHandler(c *gin.Context) {
	fmt.Println("Endpoint Hit: /api/user/")
	users := UserQueryAllField()
	for _, user := range users {
		fmt.Printf("id: %v, name: %v, password: %v\n", user.Id, user.Name, user.Password)
	}
	bytes, _ := json.Marshal(users)
	c.JSON(200, string(bytes))
}

func getUserHandler(c *gin.Context) {
	fmt.Println("Endpoint Hit: /api/user/:id")
	uid := c.Param("id")
	fmt.Println(uid)
	c.JSON(200, uid)
}
