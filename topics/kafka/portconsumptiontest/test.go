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
    topic string = "testtopic"
)

func main () {
    // guard: command line arguments
    if (len(os.Args) != 3) {
        fmt.Println("Syntax: "+os.Args[0]+" KAFKA_BROKERS GROUPNAME")
        fmt.Println("        "+os.Args[0]+" localhost:9092 8785a8e0-b6b3-42ee-a558-608d51e2a127")
        os.Exit(1)
    }
    var brokers []string = strings.Split(os.Args[1], ",")
    var group     string = os.Args[2]
    
    fmt.Printf("About to connect to %s with group id %s and listen for topic %s\n", brokers, group, topic)
    
    var i int = 1
    for {
        go func (i int) {
            reader := kafka.NewReader(kafka.ReaderConfig{
                Brokers:  brokers,
                GroupID:  group,
                Topic:    fmt.Sprintf("%s%d", topic, i),
                MinBytes: 0, // 10kB
                MaxBytes: 10e6, // 10MB
            })
            defer reader.Close()
            
            fmt.Println(fmt.Sprintf("Consumer %d started", i))
            
            for {
                m, err := reader.ReadMessage(context.Background())
                if err != nil {
                    fmt.Println(err)
                } else {
                    fmt.Println(m)
                }
            }
        }(i)
        
        time.Sleep(5 * time.Second)
        i++
    }
}
