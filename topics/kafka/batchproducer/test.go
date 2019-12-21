package main

import (
    "fmt"
    "os"
    "context"
    "time"
    
    kafka "github.com/segmentio/kafka-go"
)

const (
    id string = "testproducer"
    topic string = "testtopic"
)

func producer (channel chan string, wait chan byte, w *kafka.Writer) {
    defer close(wait)
    
    for message := range channel {
        var messages []kafka.Message = []kafka.Message{
            kafka.Message{Key: []byte("Payload"), Value: []byte(fmt.Sprintf("%s 0", message))},
            kafka.Message{Key: []byte("Payload"), Value: []byte(fmt.Sprintf("%s 1", message))},
            kafka.Message{Key: []byte("Payload"), Value: []byte(fmt.Sprintf("%s 2", message))},
            kafka.Message{Key: []byte("Payload"), Value: []byte(fmt.Sprintf("%s 3", message))},
            kafka.Message{Key: []byte("Payload"), Value: []byte(fmt.Sprintf("%s 4", message))},
            kafka.Message{Key: []byte("Payload"), Value: []byte(fmt.Sprintf("%s 5", message))},
            kafka.Message{Key: []byte("Payload"), Value: []byte(fmt.Sprintf("%s 6", message))},
            kafka.Message{Key: []byte("Payload"), Value: []byte(fmt.Sprintf("%s 7", message))},
            kafka.Message{Key: []byte("Payload"), Value: []byte(fmt.Sprintf("%s 8", message))},
            kafka.Message{Key: []byte("Payload"), Value: []byte(fmt.Sprintf("%s 9", message))},
        }
        fmt.Printf("%d new_message msg=\"%s\"\n", time.Now().UnixNano(), message)
        var err = w.WriteMessages(context.Background(), messages...)
        if err != nil {
            fmt.Println(err)
        }
        
        fmt.Printf("%d send_complete\n", time.Now().UnixNano())
    }
    
    // finish
    wait <- 0
}

func main () {
    var brokers string = os.Args[1]
    fmt.Printf("About to connect to %s\n", brokers)
    
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{brokers},
        Topic:   topic,
        Balancer: &kafka.LeastBytes{},
        BatchTimeout: 10 * time.Millisecond,
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
