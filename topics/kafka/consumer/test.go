package main

import (
    "fmt"
    "os"
    "context"
    "strings"
    "time"
    
    kafka "github.com/segmentio/kafka-go"
)

const (
    id string = "testproducer"
    topic string = "testtopic"
)

func main () {
    var brokers []string = strings.Split(os.Args[1], ",")
    var group string     = os.Args[2]
    
    fmt.Printf("About to connect to %s with group id %s and listen for topic %s\n", brokers, group, topic)
    
    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers:  brokers,
        GroupID:  group,
        Topic:    topic,
        MinBytes: 0, // 10kB
        MaxBytes: 10e6, // 10MB
    })
    defer reader.Close()
    
    for {
        fmt.Printf("%d read_begun\n", time.Now().UnixNano())
        m, err := reader.ReadMessage(context.Background())
        fmt.Printf("%d read_complete topic:%v partition:%v offset:%v key:%s value:%s\n", time.Now().UnixNano(), m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
        if err != nil {
            fmt.Println(err)
        }
    }
}
