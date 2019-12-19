package main

import (
    "fmt"
    "os"
    "context"
    
    kafka "github.com/segmentio/kafka-go"
)

const (
    id string = "testproducer"
    topic string = "testtopic"
)

func producer (channel chan string, wait chan byte, w *kafka.Writer) {
    defer close(wait)
    
    for message := range channel {
        w.WriteMessages(context.Background(), kafka.Message{
            Key:   []byte("Payload"),
            Value: []byte(message),
        })
    }
    
    // finish
    wait <- 0
}

func main () {
    var brokers string = os.Args[1]
//    var group string   = os.Args[2]
    
    fmt.Printf("About to connect to %s\n", brokers)
    
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{brokers},
        Topic:   topic,
        Balancer: &kafka.LeastBytes{},
    })
    defer writer.Close()
    
    var channel chan string = make(chan string, 10)
    var wait chan byte = make(chan byte)
    
    go producer(channel, wait, writer)
    
    for i := 0;  i<100; i++ {
        channel <- fmt.Sprintf("count: %d", i)
    }
    
    // wait for operations to finish
    close(channel)
    <- wait
}
