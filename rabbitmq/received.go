package main

import "fmt"
import "github.com/streadway/amqp"

func main(){
    conn, err := amqp.Dial("amqp://admin:admin@localhost:5672")
    if err != nil {
        fmt.Println("Connect rabbitmq error")
        return
    }
    ch, err := conn.Channel()
    if err != nil {
        fmt.Println("Open Channel error ", err)
        return
    }
    queue, err := ch.QueueDeclare(
        "hello",
        false,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        fmt.Println("Create queue error ", err)
        return
    }
    msg, err := ch.Consume(
        queue.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    forever := make(chan bool)
    go func() {
        for data := range msg {
            fmt.Println("Msg: ", string(data.Body))
        }
    }()
    fmt.Println("[*] Waiting for data")
    <- forever
}

