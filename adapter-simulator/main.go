package main

import (
	"k3s-nclink-apps/adapter-simulator/config"
	"log"
)

func main() {
	model := config.NewModel()
	model.Fetch("pi-ubt")
	log.Printf("model: %v\n", model.Def)
	model.Fetch("k8s-node-12")
	log.Printf("model: %v\n", model.Def)
}
