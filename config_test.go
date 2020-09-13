package goconfig

import (
	"errors"
	"os"
	"testing"

	"github.com/gosidekick/goconfig/goflags"
	"github.com/gosidekick/goconfig/structtag"
)

type testStruct struct {
	A int    `cfg:"A" cfgDefault:"100"`
	B string `cfg:"B" cfgDefault:"200"`
	C string
	N string `cfg:"-"`
	p string
	S testSub `cfg:"S"`
}

type testSub struct {
	A int        `cfg:"A" cfgDefault:"300"`
	B string     `cfg:"C" cfgDefault:"400"`
	S testSubSub `cfg:"S"`
}
type testSubSub struct {
	A int    `cfg:"A" cfgDefault:"500"`
	B string `cfg:"S" cfgDefault:"" cfgRequired:"true"`
}

func TestFindFileFormat(t *testing.T) {
	_, err := findFileFormat(".json")
	if err != ErrFileFormatNotDefined {
		t.Fatal(err)
	}
	Formats = []Fileformat{{Extension: ".json"}}
	_, err = findFileFormat(".json")
	if err != nil {
		t.Fatal(err)
	}
}

// -=-=-=-=-=-=-=-=-=

func mLoad(config interface{}) (err error) {
	return
}

func mPrepareHelp(config interface{}) (help string, err error) {
	return
}

// -=-=-=-=-=-=-=-=-
func eLoad(config interface{}) (err error) {
	err = errors.New("test")
	return
}

func ePrepareHelp(config interface{}) (help string, err error) {
	err = errors.New("test")
	return
}

// -=-=-=-=-=-=-=-=-

func TestParse(t *testing.T) {

	s := &testStruct{A: 1, S: testSub{A: 1, B: "2"}}
	File = "config.txt"

	Formats = []Fileformat{{Extension: ".json", Load: mLoad, PrepareHelp: mPrepareHelp}}

	err := Parse(s)
	if err != ErrFileFormatNotDefined {
		t.Fatal("Error ErrFileFormatNotDefined expected")
	}

	File = "config.json"

	Formats = []Fileformat{{Extension: ".json", Load: eLoad, PrepareHelp: mPrepareHelp}}

	err = Parse(s)
	if err == nil {
		t.Fatal("Error expected")
	}

	Formats = []Fileformat{{Extension: ".json", Load: mLoad, PrepareHelp: ePrepareHelp}}

	err = Parse(s)
	if err == nil {
		t.Fatal("Error expected")
	}

	Formats = []Fileformat{{Extension: ".json", Load: mLoad, PrepareHelp: mPrepareHelp}}

	err = os.Setenv("A", "900")
	if err != nil {
		t.Fatal(err)
	}
	err = os.Setenv("B", "TEST")
	if err != nil {
		t.Fatal(err)
	}

	Tag = ""
	err = Parse(s)
	if err != structtag.ErrUndefinedTag {
		t.Fatal("Error structtag.ErrUndefinedTag expected")
	}

	err = os.Setenv("S_S_S", "TEST")
	if err != nil {
		t.Fatal(err)
	}

	Tag = "cfg"
	err = Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("A", "900ERROR")

	goflags.Reset()
	err = Parse(s)
	if err == nil {
		t.Fatal("Error expected")
	}

	err = os.Setenv("A", "")
	if err != nil {
		t.Fatal(err)
	}

	goflags.Reset()
	err = Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	s1 := "test"
	goflags.Reset()
	err = Parse(s1)
	if err == nil {
		t.Fatal("Error expected")
	}

	goflags.Reset()
	err = Parse(&s1)
	if err == nil {
		t.Fatal("Error expected")
	}

	value := "test_file.json"
	err = os.Setenv(FileEnv, value)
	if err != nil {
		t.Fatal(err)
	}

	err = Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	if value != File {
		t.Fatal("File name could not be loaded")
	}

	value = "/var"
	err = os.Setenv(PathEnv, value)
	if err != nil {
		t.Fatal(err)
	}

	err = Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	if value != Path {
		t.Fatal("File path could not be loaded")
	}
}

func ExampleParse() {

	type config struct {
		Name  string `cfg:"Name" cfgDefault:"root"`
		Value int    `cfg:"Value" cfgDefault:"123"`
	}

	cfg := config{}

	err := Parse(&cfg)
	if err != nil {
		println(err)
	}

	println("Name:", cfg.Name, "Value:", cfg.Value)

}
