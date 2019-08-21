package main

// void SayHello(const char* str);
import "C"

func main(){
    C.SayHello(C.CString("iphone is Rubbish"))
}

