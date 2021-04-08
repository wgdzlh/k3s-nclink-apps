package main

import (
	"config-distribute/routes"
	"config-distribute/utils"
)

func main() {
	// userservice := service.Userservice{}
	// user := entity.NewUser("test1", "123456")
	// err := userservice.Create(user)
	// if err != nil {
	// 	log.Println("Error creating mongodb doc: ", err)
	// }
	// user.Name = "test2"
	// ret, err := userservice.Find(user)
	// if err != nil {
	// 	log.Fatalln("error geting some doc: ", err)
	// }
	// log.Println("test user: ", *ret)
	router := routes.InitRoute()
	host := utils.EnvVar("SERVER_HOST", "localhost")
	port := utils.EnvVar("SERVER_PORT", "8080")
	router.Run(host + ":" + port)
}
