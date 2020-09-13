package main

import (
	"encoding/json"
	"fmt"

	"github.com/gosidekick/goconfig"
)

// Declare config struct

type mongoDB struct {
	Host string `cfgRequired:"true"`
	Port int    `cfgDefault:"999"`
}

type systemUser struct {
	Name     string `json:"name" cfg:"name"`
	Password string `json:"passwd" cfg:"passwd" cfgRequired:"true"`
}

type configTest struct {
	Domain  string
	User    systemUser `json:"user" cfg:"user"`
	MongoDB mongoDB    `json:"mongo" cfg:"mongo"`
}

func main() {

	// Instance config struct
	config := configTest{}

	// Pass the struct instance pointer to the parser
	err := goconfig.Parse(&config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// it just print the config struct on the screen
	j, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	println(string(j))
}
