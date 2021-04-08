package db

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, "test", options.Client().ApplyURI("mongodb://someone:t0p-Secret@172.16.1.11:30017"))
	if err != nil {
		log.Fatal("Mongodb connect fail.", err)
	}
}
