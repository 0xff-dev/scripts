package main

import (
	"fmt"
	"os"
)

func GenerateFile() {
	file, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i := 0; i < 30; i++ {
		file.WriteString(fmt.Sprintf("%d\n", i))
	}
}
