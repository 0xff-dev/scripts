package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Art struct {
	ID      int
	Content string
	Author  string
}

var Db *sql.DB

func Posts(limit int) (arts []Art, err error) {
	//用Db直接查询的时候
	rows, err := Db.Query("select id,content,author from posts limit ?", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		art := Art{}
		err = rows.Scan(&art.ID, &art.Content, &art.Author)
		if err != nil {
			return
		}
		arts = append(arts, art)
	}
	rows.Close()
	return arts, nil
}

func GetPost(id int) (art Art, err error) {
	art = Art{}
	stmt, err := Db.Prepare("select * from posts where id=?")
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&art.ID, &art.Content, &art.Author)
	return
}

func (art *Art) Insert() (err error) {
	stmt, err := Db.Prepare("insert into posts(content,author) values(?,?)")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(art.Content, art.Author)
	return
}

func (art *Art) Update() (err error) {
	_, err = Db.Exec("update posts set content=?,author=?", art.Content, art.Author)
	return
}

func (art *Art) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id=?", art.ID)
	return
}

func init() {
	var err error
    Db, err = sql.Open("mysql", "user:passwd@/test")
	if err != nil {
		panic(err)
	}
}
func main() {
	res, _ := Posts(2)
	for _, art := range res {
		fmt.Println(art)
	}
	art1 := Art{
		ID:      2,
		Content: "Post Just",
		Author:  "Just",
	}
	err := art1.Insert()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("insert %v success!", art1)
	}
	err = art1.Delete()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("delete %v success1", art1)
	}
}
