package main

import (
	"k3s-nclink-apps/adapter-simulator/config"
	"k3s-nclink-apps/adapter-simulator/mqtt"
	"k3s-nclink-apps/utils"
	"log"
)

func main() {
	hostname := utils.GetEnvOrExit("HOSTNAME")
	model := config.NewModel()
	model.Fetch(hostname)
	log.Printf("model: %v\n", model.Def)
	// model.Fetch("k8s-node-12")
	// log.Printf("model: %v\n", model.Def)
	mqtt.Run(model)
}
