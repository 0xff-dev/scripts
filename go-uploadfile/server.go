package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func postFile(writer http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		now := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(now, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		writer.Write([]byte(token))
	} else {
		//post请求
		req.ParseMultipartForm(32 << 20)
		file, handler, err := req.FormFile("uploadfile")
		if err != nil {
			fmt.Println("Upload file error")
			return
		}
		defer file.Close()
		f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf("Open file %s error %s", handler.Filename, err)
			return
		}
		defer f.Close()
		fmt.Fprintf(writer, "%v", handler.Header)
		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/upload", postFile)
	err := http.ListenAndServe(":8080", nil)
	// 实际上要用log, 不要在屏幕输出
	if err != nil {
		fmt.Println("Listen and Serve error")
	}
}
