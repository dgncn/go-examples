package main

import (
	"dino/dynowebportal"
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Url string `json:"ServerAddress"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	config := new(Configuration)
	decoder := json.NewDecoder(file)
	decoder.Decode(config)
	dynowebportal.RunWebPortal(config.Url)
}
