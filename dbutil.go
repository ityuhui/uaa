package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var g_sqlDb *sql.DB
var sqlDberr error

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
	g_sqlDb, sqlDberr = sql.Open("mysql", dbDSN)
	if sqlDberr != nil {
		panic(sqlDberr)
	} else {
		fmt.Println("Open DB successfully.")
	}
	g_sqlDb.SetConnMaxLifetime(time.Minute * 3)
	g_sqlDb.SetMaxOpenConns(10)
	g_sqlDb.SetMaxIdleConns(10)

	sqlDberr = g_sqlDb.Ping()
	if sqlDberr != nil {
		panic(sqlDberr.Error()) // proper error handling instead of panic in your app
	} else {
		fmt.Println("Ping DB successfully.")
	}
}

func UserQueryAllField() []*User {
	users := make([]*User, 0)
	rows, err := g_sqlDb.Query("SELECT * FROM `users` limit ?", 100)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name, &user.Password)
		fmt.Printf("id:%d name:%s password:%s\n", user.Id, user.Name, user.Password)
		users = append(users, &user)
	}
	return users
}
