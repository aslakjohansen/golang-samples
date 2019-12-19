package main

import (
    "fmt"
    "os"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
    id string = "testproducer"
    topic string = "testtopic"
)

func producer (channel chan string, wait chan byte, p kafka.Producer) {
    defer close(wait)
    
    for message := range channel {
        err = p.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
            Value: []byte(message)},
        )
        if err != nil {
            fmt.Printf("Failed to transmit: %s\n", err)
            os.Exit(1)
        }
    }
    
    // finish
    wait <- 0
}

func main () {
    brokers string := os.Args[1]
    group string   := os.Args[2]
    
    fmt.Printf("About to connect to %s", prokers)
    
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": brokers,
        "group.id":          group,
        "client.id": myid,
        "default.topic.config": kafka.ConfigMap{'acks': 'all'}
    })
    if err != nil {
        fmt.Printf("Failed to create producer: %s\n", err)
        os.Exit(1)
    }
    
    channel chan string := make(chan string, 10)
    channel byte := make(chan byte)
    
    go sender(channel, wait, p)
    
    for i := 0;  i<100; i++ {
        channel <- fmt.Sprintf("%d", i)
    }
    
    // wait for operations to finish
    close(channel)
    <- wait
}
