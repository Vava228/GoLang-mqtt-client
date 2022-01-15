package main

import (
    "fmt"
    "os"
    "time"

    mqtt "github.com/eclipse/paho.mqtt.golang"
)

func messageHandler(c mqtt.Client, msg mqtt.Message) {
    fmt.Printf("MSG: %s", msg.Payload())
    fmt.Printf("from topic: \n", msg.Topic())
}

func connLostHandler(c mqtt.Client, err error) {
    fmt.Printf("Connection lost, reason: %v\n", err)
}

func main() {
    //ClientOptions
    opts := mqtt.NewClientOptions().
        AddBroker("tcp://test.mosquitto.org:1883").
        SetClientID("group-one").
        SetDefaultPublishHandler(messageHandler).
        SetConnectionLostHandler(connLostHandler)

    opts.OnConnect = func(c mqtt.Client) {
        fmt.Printf("Client connected, subscribing to: all topics\n")

        if token := c.Subscribe("/#", 0, nil); token.Wait() && token.Error() != nil {
            fmt.Println(token.Error())
            os.Exit(1)
        }
    }

    c := mqtt.NewClient(opts)
    if token := c.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    for {
        time.Sleep(50  * time.Millisecond)
    }
}
