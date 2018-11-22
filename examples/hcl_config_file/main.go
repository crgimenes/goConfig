/*
Example with configuration file.
*/
package main

import (
	"encoding/json"
	"fmt"

	"github.com/crgimenes/goconfig"
	_ "github.com/crgimenes/goconfig/hcl"
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
	DebugMode bool `hcl:"debug" cfg:"debug" cfgDefault:"false"`
	Domain    string
	User      systemUser `hcl:"user" cfg:"user"`
	MongoDB   mongoDB    `hcl:"mongodb" cfg:"mongodb"`
}

func main() {
	config := configTest{}

	goconfig.File = "config.hcl"
	err := goconfig.Parse(&config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// just print struct on screen
	j, _ := json.MarshalIndent(config, "", "\t")
	fmt.Println(string(j))
}
