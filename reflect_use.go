package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Pwd  string
}

func UseReflect(elem interface{}) []string {
	elemType := reflect.TypeOf(elem)
	// 一个结构体
	if elemType.Kind() == reflect.Struct {
		fieldNum := elemType.NumField()
		fieldNames := make([]string, fieldNum)
		for i := 0; i < fieldNum; i++ {
			fieldNames = append(fieldNames, elemType.Field(i).Name)
		}
		return fieldNames
	}
	return []string{}
}

func main() {
	for _, res := range UseReflect(User{}) {
		fmt.Print(res, " ")
	}
}
