package main

import (
	"bytes"
	"encoding/json"
	"io"
	"k3s-nclink-apps/utils"
	"log"
	"net/http"
)

type tokenResponse struct {
	Token string `json:"token"`
}

func main() {
	configUser := utils.GetEnvOrExit("CONFIG_USER")
	configPass := utils.GetEnvOrExit("CONFIG_PASS")
	configHost := utils.GetEnvOrExit("CONFIG_HOST")

	configURL := "http://" + configHost
	loginURL := configURL + "/login"
	loginBody, _ := json.Marshal(map[string]interface{}{
		"name":     configUser,
		"password": configPass,
	})
	resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(loginBody))
	if err != nil {
		log.Fatalln("login error: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("resp body read failed")
	}
	token := tokenResponse{}
	if err = json.Unmarshal(body, &token); err != nil {
		log.Fatalln("token parse error: ", err)
	}

	log.Printf("resp body: %s\ntoken: %s\n", body, token.Token)

	configClient := new(http.Client)

	pingURL := configURL + "/ping"
	req, err := http.NewRequest("GET", pingURL, nil)
	if err != nil {
		log.Fatalln("ping req compose failed: ", err)
	}
	req.Header.Add("Authorization", "Bearer "+token.Token)

	resp1, err := configClient.Do(req)
	if err != nil {
		log.Fatalln("ping failed: ", err)
	}
	defer resp1.Body.Close()
	body, err = io.ReadAll(resp1.Body)
	if err != nil {
		log.Fatalln("ping body parse failed: ", err)
	}
	log.Printf("\nping status code: %d\n body: %s\n", resp1.StatusCode, body)
}
