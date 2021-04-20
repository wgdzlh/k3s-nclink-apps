package db

import (
	"fmt"
	"k3s-nclink-apps/utils"
	"log"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	// Setup the mgm default config
	mongoUser := utils.GetEnvOrExit("MONGO_USER")
	mongoPass := utils.GetEnvOrExit("MONGO_PASS")
	mongoHost := utils.GetEnvOrExit("MONGO_ADDR")
	mongoDB := utils.GetEnvOrExit("MONGO_DB")
	mongoURL := fmt.Sprintf("mongodb://%s:%s@%s", mongoUser, mongoPass, mongoHost)

	opts := options.Client().ApplyURI(mongoURL)

	err := mgm.SetDefaultConfig(nil, mongoDB, opts)
	if err != nil {
		log.Fatalf("Mongodb set config fail: %v", err)
	}
	_, client, _, err := mgm.DefaultConfigs()
	if err != nil {
		log.Fatalf("Mongodb get client fail: %v", err)
	}
	if err = client.Ping(mgm.NewCtx(10*time.Second), readpref.Primary()); err != nil {
		log.Fatalf("Mongodb connect err: %v", err)
	}
}
