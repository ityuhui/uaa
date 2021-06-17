package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initUserAPIEndpoint(r *gin.Engine) {
	r.POST("/api/user", createUserHandler)
	r.GET("/api/user", listUserHandler)
	r.GET("/api/user/:id", getUserHandler)
	r.DELETE("/api/user/:id", deleteUserHandler)
	printUserEndpoint()
}

func printUserEndpoint() {
	fmt.Print(`TODO
PUT /api/user/{id} Update
-----------------
`)
}

func createUserHandler(c *gin.Context) {
	fmt.Println("Endpoint Hit: POST /api/user/")
	name := c.PostForm("name")
	if name != "" {
		c.JSON(http.StatusOK, "Create user <"+name+"> successfully.")
	} else {
		c.JSON(http.StatusBadRequest, "Fail to create user.")
	}
}

func listUserHandler(c *gin.Context) {
	fmt.Println("Endpoint Hit: GET /api/user/")
	users := UserQueryAllField()
	for _, user := range users {
		fmt.Printf("id: %v, name: %v, password: %v\n", user.Id, user.Name, user.Password)
	}
	bytes, _ := json.Marshal(users)
	c.JSON(200, string(bytes))
}

func getUserHandler(c *gin.Context) {
	fmt.Println("Endpoint Hit: GET /api/user/:id")
	uid := c.Param("id")
	fmt.Println(uid)
	c.JSON(200, uid)
}

func deleteUserHandler(c *gin.Context) {
	fmt.Println("Endpoint Hit: DELETE /api/user/:id")
	uid := c.Param("id")
	err := deleteUser(uid)
	if err != nil {
		c.JSON(400, "cannot delete the user "+uid)
	} else {
		c.JSON(200, uid)
	}
}
