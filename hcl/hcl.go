package hcl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gosidekick/goconfig"
	"github.com/fatih/structs"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
)

func init() {
	f := goconfig.Fileformat{
		Extension:   ".hcl",
		Load:        LoadHCL,
		PrepareHelp: PrepareHelp,
	}
	goconfig.Formats = append(goconfig.Formats, f)
}

// LoadHCL config file
func LoadHCL(config interface{}) (err error) {
	configFile := filepath.Join(goconfig.Path, goconfig.File)
	_, err = os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) && !goconfig.FileRequired {
			err = nil
		}
		return
	}
	byt, err := ioutil.ReadFile(configFile) // nolint
	if err != nil && err != io.EOF {
		return
	}
	err = hcl.Unmarshal(byt, config)
	return
}

// PrepareHelp return help string for this file format.
func PrepareHelp(config interface{}) (help string, err error) {
	structs.DefaultTagName = "hcl"
	m := structs.Map(config)
	byt, err := json.Marshal(m)
	if err != nil {
		return
	}
	ast, err := jsonParser.Parse(byt)
	if err != nil {
		return
	}
	buff := &bytes.Buffer{}
	err = printer.Fprint(buff, ast)
	if err != nil {
		return
	}
	help = fmt.Sprintf("\n '=' BEFORE '{' IS OPTIONAL\n\n %s", buff.String())
	return
}
