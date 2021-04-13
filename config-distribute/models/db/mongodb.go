package db

import (
	"fmt"
	"k3s-nclink-apps/utils"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// Setup the mgm default config
	mongoUser := utils.GetEnvOrExit("MONGO_USER")
	mongoPass := utils.GetEnvOrExit("MONGO_PASS")
	mongoHost := utils.GetEnvOrExit("MONGO_HOST")
	mongoURL := fmt.Sprintf("mongodb://%s:%s@%s", mongoUser, mongoPass, mongoHost)

	err := mgm.SetDefaultConfig(nil, "test", options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal("Mongodb connect fail.", err)
	}
}
