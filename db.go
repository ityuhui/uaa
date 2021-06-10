package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var g_sqlDb *sql.DB

const (
	USER_NAME = "uaa"
	PASS_WORD = "password"
	HOST      = "localhost"
	PORT      = "30306"
	DATABASE  = "uaa"
	CHARSET   = "utf8"
)

func connectDB() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
	g_sqlDb, err := sql.Open("mysql", dbDSN)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Open DB successfully.")
	}
	defer g_sqlDb.Close()
	g_sqlDb.SetConnMaxLifetime(time.Minute * 3)
	g_sqlDb.SetMaxOpenConns(10)
	g_sqlDb.SetMaxIdleConns(10)

	err = g_sqlDb.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	} else {
		fmt.Println("Ping DB successfully.")
	}
}

func StructQueryAllField() []*User {

	// 通过切片存储
	users := make([]*User, 0)
	fmt.Println("111")
	rows, _ := g_sqlDb.Query("SELECT * FROM `users` limit ?", 100)
	fmt.Println("222")
	// 遍历
	var user User
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Password)
		users = append(users, &user)
	}
	fmt.Println("333")
	return users
}
