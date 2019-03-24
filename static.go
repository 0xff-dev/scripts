// 获取项目目录
package main

import (
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	pwd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	mux.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir(pwd))))
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err.Error())
	}
}
