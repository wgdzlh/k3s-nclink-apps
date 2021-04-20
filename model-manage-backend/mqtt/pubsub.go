package mqtt

import (
	"fmt"
	"k3s-nclink-apps/utils"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	resetModelPubFormat = "nclink/resetmodel/request/%s"
	// resetModelSubFormat = "nclink/resetmodel/response/%s"
)

var (
	client mqtt.Client
	// hostname string
	// resetModelSub string
)

func connectHandler(client mqtt.Client) {
	log.Println("MQTT Connected")
}

func connectionLostHandler(client mqtt.Client, err error) {
	log.Printf("MQTT Connection Lost: %v\n", err)
}

// func onResetModel(client mqtt.Client, msg mqtt.Message) {
// }

func setup() *mqtt.ClientOptions {
	broker := utils.GetEnvOrExit("MQTT_ADDR")
	user := utils.GetEnvOrExit("MQTT_USER")
	pass := utils.GetEnvOrExit("MQTT_PASS")
	clientId := utils.EnvVar("MQTT_CLIENT", "go_mqtt_model_manager")

	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID(clientId)
	options.SetUsername(user)
	options.SetPassword(pass)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler
	return options
}

func ResetModel(adapterName string) {
	var token mqtt.Token
	if client == nil {
		// hostname = adapterName
		client = mqtt.NewClient(setup())
		token = client.Connect()
		if token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	// else {
	// client.Unsubscribe(resetModelSub).Wait()
	// }
	// resetModelSub = fmt.Sprintf(resetModelSubFormat, hostname)
	// client.Subscribe(resetModelSub, 1, onResetModel).Wait()

	topic := fmt.Sprintf(resetModelPubFormat, adapterName)
	token = client.Publish(topic, 2, false, "")
	log.Printf("Model reseted on host '%s'\n", adapterName)
	token.Wait()
}
