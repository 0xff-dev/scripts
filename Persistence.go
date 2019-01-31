package main

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func useGob(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0660)
	if err != nil {
		panic(err)
	}

	// 读取数据, 二进制
	var res Post
	datas, _ := ioutil.ReadFile(filename)
	buffer = bytes.NewBuffer(datas)
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(&res)
	if err != nil {
		panic(err)
	}
	fmt.Println("res: ", res)
}

func storeIntoCsv(posts []Post) {
	file, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	csvCwriter := csv.NewWriter(file)
	for _, post := range posts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err = csvCwriter.Write(line)
		if err != nil {
			panic(err)
		}
	}
	csvCwriter.Flush()
	file, err = os.Open("data.csv")
	defer file.Close()
	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1
	records, _ := csvReader.ReadAll()
	for _, line := range records {
		id, _ := strconv.ParseInt(line[0], 0, 0)
		post := Post{
			Id:      int(id),
			Content: line[1],
			Author:  line[2],
		}
		fmt.Printf("%v \n", post)
	}
}

func main() {
	posts := []Post{
		{Id: 1, Content: "Post1", Author: "Coco"},
		{Id: 2, Content: "Post2", Author: "Coco"},
		{Id: 3, Content: "Post43", Author: "Just"},
		{Id: 4, Content: "Post4", Author: "Justzs"},
	}
	storeIntoCsv(posts)
	useGob(Post{Id: 5, Content: "Post5", Author: "Coco"}, "testBin")
}
