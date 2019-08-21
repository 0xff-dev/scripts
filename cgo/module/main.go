package main

//#include<hello.h>
import "C"

func main(){
    C.Hello(C.CString("Iphone is Rubbish"))
}

