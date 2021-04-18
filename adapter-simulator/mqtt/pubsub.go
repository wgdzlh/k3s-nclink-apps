package mqtt

import (
	"fmt"
	"k3s-nclink-apps/adapter-simulator/config"
	pb "k3s-nclink-apps/configmodel"
	"k3s-nclink-apps/utils"
	"log"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	samplePubFormat = "nclink/sample/%s/%s"
	sampleMsgFormat = "{\"value\": %v}"
	querySubFormat  = "nclink/query/request/%s/%s"
	queryPubFormat  = "nclink/query/response/%s/%s"
	tweakSubFormat  = "nclink/tweak/request/%s/%s"
	tweakPubFormat  = "nclink/tweak/response/%s/%s"
)

var (
	options    *mqtt.ClientOptions
	devId      string
	samples    []*pb.Sample
	samplePubs []string
	querySubs  []string
	tweakSubs  []string
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

func processSamplePubs() {
	samplePubs = make([]string, len(samples))
	for i, q := range samples {
		samplePubs[i] = fmt.Sprintf(samplePubFormat, devId, q.Sensor)
	}
}

func processQuerySubs(querys []*pb.Query) {
	querySubs = make([]string, len(querys))
	for i, q := range querys {
		querySubs[i] = fmt.Sprintf(querySubFormat, devId, q.Sensor)
	}
}

func processTweakSubs(tweaks []*pb.Tweak) {
	tweakSubs = make([]string, len(tweaks))
	for i, t := range tweaks {
		tweakSubs[i] = fmt.Sprintf(tweakSubFormat, devId, t.Register)
	}
}

func setup(model *config.Model) {
	rand.Seed(time.Now().UnixNano())
	broker := utils.GetEnvOrExit("MQTT_ADDR")
	user := utils.GetEnvOrExit("MQTT_USER")
	pass := utils.GetEnvOrExit("MQTT_PASS")
	clientId := utils.EnvVar("MQTT_CLIENT", "go_mqtt_example")

	options = mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID(clientId)
	options.SetUsername(user)
	options.SetPassword(pass)
	options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	devId = model.DevId
	modelDef := model.Def
	samples = modelDef.Sample
	processSamplePubs()
	processQuerySubs(modelDef.Query)
	processTweakSubs(modelDef.Tweak)
}

func Run(model *config.Model) {
	setup(model)
	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for _, topic := range querySubs {
		token = client.Subscribe(topic, 1, nil)
		token.Wait()
		log.Printf("Subscribed to query topic %s\n", topic)
	}

	for _, topic := range tweakSubs {
		token = client.Subscribe(topic, 1, nil)
		token.Wait()
		log.Printf("Subscribed to tweak topic %s\n", topic)
	}

	for i, topic := range samplePubs {
		sensor := samples[i].Sensor
		interval := time.Duration(1000 / samples[i].Rate)
		ticker := time.NewTicker(interval * time.Millisecond)
		go func(topic, sensor string) {
			for {
				<-ticker.C
				value := rand.Intn(100)
				text := fmt.Sprintf(sampleMsgFormat, value)
				token = client.Publish(topic, 0, false, text)
				log.Printf("Sample %s: %v", sensor, value)
				token.Wait()
			}
		}(topic, sensor)
	}

	time.Sleep(24 * 36500 * time.Hour)
	client.Disconnect(100)
	log.Println("Disonnected")
}
