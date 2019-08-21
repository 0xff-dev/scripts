package main

import "C"
import "fmt"

//export Hello
func Hello(str *C.char){
    fmt.Println(C.GoString(str))
}

