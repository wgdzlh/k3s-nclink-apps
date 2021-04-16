package mqtt

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func messagePubHandler(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Message %s received on topic %s\n", msg.Payload(), msg.Topic())
}

func connectHandler(client mqtt.Client) {
	log.Println("Connected")
}

func connectionLostHandler(client mqtt.Client, err error) {
	log.Printf("Connection Lost: %v\n", err)
}

func run() {
	var broker = "tcp://172.16.1.10:2883"
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID("go_mqtt_example")
	options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topic := "topic/secret"
	token = client.Subscribe(topic, 1, nil)
	token.Wait()
	log.Printf("Subscribed to topic %s\n", topic)

	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("message - %d", i)
		token = client.Publish(topic, 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}

	client.Disconnect(100)
}
