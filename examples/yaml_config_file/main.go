/*
Example with configuration file.
*/
package main

import (
	"encoding/json"

	"github.com/gosidekick/goconfig"
	_ "github.com/gosidekick/goconfig/yaml"
)

type mongoDB struct {
	Host string `yaml:"host" cfg:"host" cfgDefault:"example.com"`
	Port int    `yaml:"port" cfg:"port" cfgDefault:"999"`
}

type systemUser struct {
	Name     string `yaml:"name" cfg:"name"`
	Password string `yaml:"passwd" cfg:"passwd"`
}

type configTest struct {
	DebugMode bool       `yaml:"debug" cfg:"debug" cfgDefault:"false"`
	Domain    string     `yaml:"domain"`
	User      systemUser `yaml:"user" cfg:"user"`
	MongoDB   mongoDB    `yaml:"mongodb" cfg:"mongodb"`
}

func main() {
	config := configTest{}

	goconfig.File = "config.yaml"
	err := goconfig.Parse(&config)
	if err != nil {
		println(err)
		return
	}

	// just print struct on screen
	j, _ := json.MarshalIndent(config, "", "\t")
	println(string(j))
}
