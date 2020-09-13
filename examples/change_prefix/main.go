package main

import (
	"encoding/json"

	"github.com/gosidekick/goconfig"
)

// Declare config struct

type mongoDB struct {
	Host string `cfgDefault:"example.com"`
	Port int    `cfgDefault:"999"`
}

type systemUser struct {
	Name     string `json:"name" cfg:"name"`
	Password string `json:"passwd" cfg:"passwd"`
}

type configTest struct {
	Domain  string
	User    systemUser `json:"user" cfg:"user"`
	MongoDB mongoDB    `json:"mongo" cfg:"mongo"`
}

func main() {

	// Instance config struct
	config := configTest{}

	// Adds a prefix to the environment variables.
	goconfig.PrefixEnv = "EXAMPLE"

	// Pass the struct instance pointer to the parser
	err := goconfig.Parse(&config)
	if err != nil {
		println(err)
		return
	}

	// it just print the config struct on the screen
	j, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		println(err)
		return
	}
	println(string(j))
}
