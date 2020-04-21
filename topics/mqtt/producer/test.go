package main

import (
    "fmt"
    "time"
    
    "github.com/eclipse/paho.mqtt.golang"
)

var (
    broker string = "tcp://192.168.1.33:1883"
)

func main () {
    options := mqtt.NewClientOptions()
    options.AddBroker(broker)
    
    client := mqtt.NewClient(options)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
    
    var i int = 0
    var message string;
    for {
        message = fmt.Sprintf("%d", i);
        client.Publish("test/counter", 1, false, message)
        
        time.Sleep(100 * time.Millisecond)
        
        i++
    }
}

