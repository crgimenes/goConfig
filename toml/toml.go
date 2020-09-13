package toml

import (
	"os"
	"path/filepath"
	"reflect"

	"github.com/gosidekick/goconfig"
	"github.com/pelletier/go-toml"
)

func init() {
	f := goconfig.Fileformat{
		Extension:   ".toml",
		Load:        LoadTOML,
		PrepareHelp: PrepareHelp,
	}
	goconfig.Formats = append(goconfig.Formats, f)
}

// LoadTOML config file
func LoadTOML(config interface{}) (err error) {
	configFile := filepath.Join(goconfig.Path, goconfig.File)
	_, err = os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) && !goconfig.FileRequired {
			err = nil
		}
		return
	}
	var tree *toml.Tree
	tree, err = toml.LoadFile(configFile)
	if err != nil {
		return
	}
	err = tree.Unmarshal(config)
	return
}

// PrepareHelp return help string for this file format.
func PrepareHelp(config interface{}) (help string, err error) {
	var byt []byte
	cfg := reflect.ValueOf(config).Elem()
	byt, err = toml.Marshal(cfg)
	if err != nil {
		return
	}
	help = string(byt)
	return
}
