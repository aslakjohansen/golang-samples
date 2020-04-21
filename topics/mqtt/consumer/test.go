package main

import (
    "fmt"
    
    "github.com/eclipse/paho.mqtt.golang"
)

var (
    broker string = "tcp://192.168.1.33:1883"
)

func callback (client mqtt.Client, message mqtt.Message) {
    var topic string = message.Topic()
    var payload string = string(message.Payload())
    fmt.Println(topic, " -> ", payload)
}


func main () {
    options := mqtt.NewClientOptions()
    options.AddBroker(broker)
    
    client := mqtt.NewClient(options)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
    
    if token := client.Subscribe("test/counter", 2, callback); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
    
    // block forever
    select {}
}

