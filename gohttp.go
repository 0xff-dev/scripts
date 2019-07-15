package main

import (
	"fmt"
	"net/http"
	"time"
)

// Client 控制http的请求头，需要自己创建client
func Client() {
	client := &http.Client{
		CheckRedirect: rdirectFunc,
	}
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		fmt.Errorf("create new req error: %s", err)
		return
	}
	req.Header.Add("If-None-Match", `w/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("client request error %s", err)
		return
	}
	defer resp.Body.Close()
}

// Transport 代理，TLS， keep-alive等配置
func Transport() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{
		Transport: tr,
	}
}
func main() {
	Client()
}
