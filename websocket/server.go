package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var rep string
		if err = websocket.Message.Receive(ws, &rep); err != nil {
			fmt.Println("Cann't received")
			break
		}
		fmt.Println("Received from client: ", rep)
		msg := "RRRRR: " + rep
		fmt.Println("send to client: " + msg)
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Cann't send")
			break
		}
	}
}
func main() {
	http.Handle("/", websocket.Handler(Echo))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("err")
	}
}
