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

// type tokenResponse struct {
// 	Token string `json:"token"`
// }

// func main() {
// configUser := utils.GetEnvOrExit("CONFIG_USER")
// configPass := utils.GetEnvOrExit("CONFIG_PASS")
// configHost := utils.GetEnvOrExit("CONFIG_HOST")
// configPort := utils.GetEnvOrExit("CONFIG_PORT")

// configURL := "http://" + configHost + ":" + configPort
// loginURL := configURL + "/login"
// loginBody, _ := json.Marshal(map[string]interface{}{
// 	"name":     configUser,
// 	"password": configPass,
// })
// resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(loginBody))
// if err != nil {
// 	log.Fatalln("login error: ", err)
// }
// defer resp.Body.Close()

// body, err := io.ReadAll(resp.Body)
// if err != nil {
// 	log.Fatalln("resp body read failed")
// }
// token := tokenResponse{}
// if err = json.Unmarshal(body, &token); err != nil {
// 	log.Fatalln("token parse error: ", err)
// }

// log.Printf("resp body: %s\ntoken: %s\n", body, token.Token)

// configClient := new(http.Client)

// pingURL := configURL + "/ping"
// req, err := http.NewRequest("GET", pingURL, nil)
// if err != nil {
// 	log.Fatalln("ping req compose failed: ", err)
// }
// req.Header.Add("Authorization", "Bearer "+token.Token)

// resp1, err := configClient.Do(req)
// if err != nil {
// 	log.Fatalln("ping failed: ", err)
// }
// defer resp1.Body.Close()
// body, err = io.ReadAll(resp1.Body)
// if err != nil {
// 	log.Fatalln("ping body parse failed: ", err)
// }
// log.Printf("\nping status code: %d\n body: %s\n", resp1.StatusCode, body)
// }
