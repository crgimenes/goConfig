package main

import (
	"log"

	"github.com/crgimenes/goconfig"
	_ "github.com/crgimenes/goconfig/json"
)

/*
step 1: Declare your configuration struct,
it may or may not contain substructures.
*/

type mongoDB struct {
	Host string `cfgDefault:"example.com" cfgHelper:"mongoDB host URL"`
	Port int    `cfgDefault:"999" cfgHelper:"mongoDB http port"`
}

type configTest struct {
	Domain    string
	DebugMode bool `json:"db" cfg:"db" cfgDefault:"false"`
	MongoDB   mongoDB
}

func main() {

	// step 2: Instantiate your structure.
	config := configTest{}

	goconfig.FileRequired = true
	goconfig.WatchConfigFile = true
	goconfig.Path = "./"
	goconfig.File = "config.json"

	// step 3: Pass the instance pointer to the parser
	updatesCh, errCh, err := goconfig.ParseAndWatch(&config)
	if err != nil {
		println(err.Error())
		return
	}

	/*
	   The parser populated your struct with the data
	   it took from environment variables and command
	   line and now you can use it.
	*/

	println("config.Domain......:", config.Domain)
	println("config.DebugMode...:", config.DebugMode)
	println("config.MongoDB.Host:", config.MongoDB.Host)
	println("config.MongoDB.Port:", config.MongoDB.Port)

	/*
		Now every time the config.json file change its content an timestamp will be
		returned on the updates channel
	*/
	for {
		select {
		case v := <-updatesCh:
			println("Updated at: ", v)
			println("config.Domain......:", config.Domain)
			println("config.DebugMode...:", config.DebugMode)
			println("config.MongoDB.Host:", config.MongoDB.Host)
			println("config.MongoDB.Port:", config.MongoDB.Port)
		case err := <-errCh:
			log.Fatal(err)
		}
	}
}
