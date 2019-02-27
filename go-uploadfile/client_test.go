package main

import "testing"

func TestUpload(t *testing.T) {
	UploadFile("test.txt", "http://127.0.0.1:8080/upload")
}
