package main

// #include<stdio.h>
/*
static void SayHello(const char* str) {
    puts(str);
}
*/
import "C"


func main(){
    C.puts(C.CString("Hello CGO!\n"))
    C.SayHello(C.CString("zshuang"))
}

