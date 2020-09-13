package ini

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/gosidekick/goconfig"
	ini "gopkg.in/ini.v1"
)

func init() {
	f := goconfig.Fileformat{
		Extension:   ".ini",
		Load:        LoadINI,
		PrepareHelp: PrepareHelp,
	}
	goconfig.Formats = append(goconfig.Formats, f)
}

// LoadINI config file
func LoadINI(config interface{}) (err error) {
	configFile := filepath.Join(goconfig.Path, goconfig.File)
	file, err := os.Open(configFile)
	if os.IsNotExist(err) && !goconfig.FileRequired {
		err = nil
		return
	} else if err != nil {
		return
	}

	err = ini.MapTo(config, file)
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
