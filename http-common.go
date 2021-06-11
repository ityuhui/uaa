package main

import (
	"fmt"
	"net/http"
)

func startAPIServer() {
	fmt.Println("REST API server: http://0.0.0.0:8080/api")
	rtr := mux.NewRouter()
	initStaticFileServer()
	initUserAPIEndpoint()
	http.ListenAndServe(":8080", nil)
}

func initStaticFileServer() {
	http.Handle("/", http.FileServer(http.Dir("public/")))
}

// AllowCrossDomain : Add http response header to allow cross domain
func AllowCrossDomain(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
	w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
	w.Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json
}
