package main

import (
	"k3s-nclink-apps/model-manage-backend/routes"
	"k3s-nclink-apps/utils"
	"log"
)

func main() {
	// userservice := service.UserService{}
	// user := entity.NewUser("admin", "123456")
	// err := userservice.Create(user)
	// if err != nil {
	// 	log.Println("Error creating mongodb doc: ", err)
	// }
	host := utils.EnvVar("SERVER_HOST", "localhost")
	port := utils.EnvVar("SERVER_PORT", "8000")
	addr := host + ":" + port
	log.Printf("start serving on: %s", addr)
	router := routes.InitRoute()
	router.Run(addr)
}
