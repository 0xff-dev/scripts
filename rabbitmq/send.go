package main

import "fmt"
import "github.com/streadway/amqp"

func main(){
    conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
    if err != nil {
        fmt.Println("Connecton refuse")
        return
    }
    defer conn.Close()
    ch, err := conn.Channel()
    if err != nil {
        fmt.Println("Create channel error ", err)
        return
    }
    queue, err := ch.QueueDeclare(
        "hello",
        false, // durable
        false, // delte when used
        false, // exclusive
        false, // no-wait
        nil, // args
    )
    if err != nil {
        fmt.Println("Create queue error ", err)
        return
    }
    data := "Hello RabbitMQ"
    err = ch.Publish(
        "",
        queue.Name, // routine key
        false, // mandatory
        false, // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body: []byte(data),
        })
    if err != nil {
        fmt.Println("Send msg error, ", err)
    }
}
