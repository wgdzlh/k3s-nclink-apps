package main

import (
	"k3s-nclink-apps/adapter-simulator/config"
	"log"
)

func main() {
	model := config.NewModel()
	log.Printf("model: %v\n", model.Fetch("pi-ubt"))
	log.Printf("model: %v\n", model.Fetch("k8s-node-12"))
}
