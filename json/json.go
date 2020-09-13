package json

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/gosidekick/goconfig"
	"github.com/gosidekick/goconfig/helper"
)

func init() {
	f := goconfig.Fileformat{
		Extension:   ".json",
		Load:        LoadJSON,
		PrepareHelp: PrepareHelp,
	}
	goconfig.Formats = append(goconfig.Formats, f)
}

// LoadJSON config file
func LoadJSON(config interface{}) (err error) {
	configFile := filepath.Join(goconfig.Path, goconfig.File)
	file, err := os.Open(configFile)
	if err != nil {
		if os.IsNotExist(err) && !goconfig.FileRequired {
			err = nil
		}
		return
	}
	defer helper.Closer(file)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return
	}

	return
}

// PrepareHelp return help string for this file format.
func PrepareHelp(config interface{}) (help string, err error) {
	var helpAux []byte
	helpAux, err = json.MarshalIndent(&config, "", "    ")
	if err != nil {
		return
	}
	help = string(helpAux)
	return
}
