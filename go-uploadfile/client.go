package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadFile(fileName, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	// 构造handler使用应该是
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", fileName)
	if err != nil {
		fmt.Println("constructer error")
		return err
	}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Open file error")
		return err
	}
	defer file.Close()
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		fmt.Println("Copy file error")
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		fmt.Println("Post file error")
		return err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(string(respBody))
	return nil
}
