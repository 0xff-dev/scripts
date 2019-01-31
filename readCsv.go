package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
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
}
