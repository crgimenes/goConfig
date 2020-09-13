/*
Example with configuration file.
*/
package main

import (
	"encoding/json"

	"github.com/gosidekick/goconfig"
	_ "github.com/gosidekick/goconfig/json"
)

type mongoDB struct {
	Host string `json:"host" cfg:"host" cfgDefault:"example.com"`
	Port int    `json:"port" cfg:"port" cfgDefault:"999"`
}

type systemUser struct {
	Name     string `json:"name" cfg:"name"`
	Password string `json:"passwd" cfg:"passwd"`
}

type configTest struct {
	DebugMode bool `json:"debug" cfg:"debug" cfgDefault:"false"`
	Domain    string
	User      systemUser `json:"user" cfg:"user"`
	MongoDB   mongoDB    `json:"mongodb" cfg:"mongodb"`
}

func main() {
	config := configTest{}

	goconfig.File = "config.json"
	err := goconfig.Parse(&config)
	if err != nil {
		println(err)
		return
	}

	// just print struct on screen
	j, _ := json.MarshalIndent(config, "", "\t")
	println(string(j))
}
