package main

import (
    "fmt"
    "os"
    "context"
    "strings"
    
    kafka "github.com/segmentio/kafka-go"
)

const (
    id string = "testproducer"
    topic string = "testtopic"
)

func main () {
    var brokers []string = strings.Split(os.Args[1], ",")
    var group string     = os.Args[2]
    
    fmt.Printf("About to connect to %s\n", brokers)
    
    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers:  brokers,
        GroupID:  group,
        Topic:    topic,
        MinBytes: 10e3, // 10kB
        MaxBytes: 10e6, // 10MB
    })
    defer reader.Close()
    
    for {
        m, err := reader.ReadMessage(context.Background())
        if err != nil {
            fmt.Println(err)
        }
        fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
    }
}
