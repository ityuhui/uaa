package main

import (
	"fmt"
	"net/http"
	_ "time"
)

// AllowCrossDomain : Add http response header to allow cross domain
func AllowCrossDomain(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
	w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
	w.Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	AllowCrossDomain(w)
	fmt.Fprintf(w, "Welcome to the /api/welcome !")
	fmt.Println("Endpoint Hit: /api/welcome")
}

func connectDB() bool {

	if err != nil {
		panic("failed to connect database")
	} else {

	}
}

func main() {
	fmt.Println("Webserver-go starts...")
	fmt.Println(" * static file serves at http://0.0.0.0:8080/")
	fmt.Println(" * RESTful API serves at http://0.0.0.0:8080/api/")
	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/api/welcome", welcomeHandler)

	http.ListenAndServe(":8080", nil)
}
