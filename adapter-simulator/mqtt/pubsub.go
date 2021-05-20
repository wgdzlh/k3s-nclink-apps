package mqtt

import (
	"fmt"
	"k3s-nclink-apps/adapter-simulator/config"
	pb "k3s-nclink-apps/configmodel"
	"k3s-nclink-apps/utils"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	samplePubFormat     = "nclink/sample/%s/%s"
	sampleMsgFormat     = "{\"sensor\":\"%s\",\"sample_value\": %v}"
	querySubFormat      = "nclink/query/request/%s/%s"
	queryPubFormat      = "nclink/query/response/%s/%s"
	queryMsgFormat      = "{\"query_value\": %v}"
	tweakSubFormat      = "nclink/tweak/request/%s/%s"
	tweakPubFormat      = "nclink/tweak/response/%s/%s"
	tweakMsgFormat      = "{\"tweaked_to\": %v}"
	resetModelSubFormat = "nclink/resetmodel/request/%s"
	resetModelPubFormat = "nclink/resetmodel/response/%s"
	retMsgFormat        = "{\"ret_msg\": \"%s\"}"
)

var (
	model         *config.Model
	client        mqtt.Client
	hostname      string
	devId         string
	samples       []*pb.Sample
	samplePubs    []string
	querySubs     []string
	tweakSubs     []string
	resetModelSub string
	endSample     chan bool
)

func messageDefaultHandler(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Default handler: message '%s' received on topic %s\n",
		msg.Payload(), msg.Topic())
}

func connectHandler(client mqtt.Client) {
	log.Println("Connected")
}

func connectionLostHandler(client mqtt.Client, err error) {
	log.Printf("Connection Lost: %v\n", err)
}

func onQueryMessage(client mqtt.Client, msg mqtt.Message) {
	sensor := strings.SplitN(msg.Topic(), "/", 5)[4]
	topic := fmt.Sprintf(queryPubFormat, devId, sensor)
	value := rand.Intn(100)
	text := fmt.Sprintf(queryMsgFormat, value)
	token := client.Publish(topic, 1, false, text)
	log.Printf("Query %s: %v\n", sensor, value)
	token.Wait()
}

func onResetModel(client mqtt.Client, msg mqtt.Message) {
	endSample <- true
	model.Fetch(hostname)
	log.Printf("Reseted model: %v\n", model.Def)
	Run(nil)
	topic := fmt.Sprintf(resetModelPubFormat, hostname)
	text := fmt.Sprintf(retMsgFormat, "reset model completed")
	token := client.Publish(topic, 0, false, text)
	log.Printf("Model reseted on host '%s'\n", hostname)
	token.Wait()
}

func onTweakMessage(client mqtt.Client, msg mqtt.Message) {
	reg := strings.SplitN(msg.Topic(), "/", 5)[4]
	in := string(msg.Payload())
	value, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		valuef, err := strconv.ParseFloat(in, 64)
		if err != nil {
			log.Printf("failed tweaking register '%s' to '%s'\n", reg, in)
			return
		}
		tweakRegister(client, reg, valuef)
	} else {
		tweakRegister(client, reg, value)
	}
}

func tweakRegister(client mqtt.Client, reg string, value interface{}) {
	topic := fmt.Sprintf(tweakPubFormat, devId, reg)
	text := fmt.Sprintf(tweakMsgFormat, value)
	token := client.Publish(topic, 0, false, text)
	log.Printf("Tweak '%s': %v\n", reg, value)
	token.Wait()
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

func setModel(model *config.Model) {
	devId = model.DevId
	modelDef := model.Def
	samples = modelDef.Sample
	processSamplePubs()
	processQuerySubs(modelDef.Query)
	processTweakSubs(modelDef.Tweak)
}

func setup() *mqtt.ClientOptions {
	rand.Seed(time.Now().UnixNano())
	hostname = utils.GetEnvOrExit("HOSTNAME")
	broker := utils.GetEnvOrExit("MQTT_ADDR")
	user := utils.GetEnvOrExit("MQTT_USER")
	pass := utils.GetEnvOrExit("MQTT_PASS")
	clientId := utils.EnvVar("MQTT_CLIENT", "go_mqtt_"+hostname)
	resetModelSub = fmt.Sprintf(resetModelSubFormat, hostname)
	endSample = make(chan bool)

	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID(clientId)
	options.SetUsername(user)
	options.SetPassword(pass)
	options.SetDefaultPublishHandler(messageDefaultHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler
	return options
}

func Run(inModel *config.Model) {
	var token mqtt.Token
	if model == nil {
		model = inModel
		client = mqtt.NewClient(setup())
		token = client.Connect()
		if token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
		client.Subscribe(resetModelSub, 2, onResetModel).Wait()
	} else {
		log.Println("restart mqtt subpub.")
		if len(querySubs) > 0 {
			client.Unsubscribe(querySubs...).Wait()
		}
		if len(tweakSubs) > 0 {
			client.Unsubscribe(tweakSubs...).Wait()
		}
	}
	setModel(model)

	for _, topic := range querySubs {
		token = client.Subscribe(topic, 1, onQueryMessage)
		token.Wait()
		log.Printf("Subscribed to query topic %s\n", topic)
	}

	for _, topic := range tweakSubs {
		token = client.Subscribe(topic, 2, onTweakMessage)
		token.Wait()
		log.Printf("Subscribed to tweak topic %s\n", topic)
	}

	for i, topic := range samplePubs {
		sensor := samples[i].Sensor
		interval := time.Duration(1000. / samples[i].Rate)
		ticker := time.NewTicker(interval * time.Millisecond)
		go func(topic string) {
			for {
				select {
				case <-endSample:
					ticker.Stop()
					return
				case <-ticker.C:
					value := rand.Intn(100)
					text := fmt.Sprintf(sampleMsgFormat, sensor, value)
					token = client.Publish(topic, 0, false, text)
					// log.Printf("Sample %s: %v\n", sensor, value)
					token.Wait()
				}
			}
		}(topic)
	}
}
