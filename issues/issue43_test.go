package issue43

import (
	"encoding/json"
	"testing"

	"github.com/gosidekick/goconfig"
	_ "github.com/gosidekick/goconfig/ini"
	_ "github.com/gosidekick/goconfig/json"
)

type boo struct {
	Boo11 bool `json:"BOO11" ini:"BOO11" cfg:"BOO11"`
	Boo12 bool `json:"BOO12" ini:"BOO12" cfg:"BOO12"`
	Boo13 bool `json:"BOO13" ini:"BOO13" cfg:"BOO13"`

	Boo21 bool `json:"BOO21" ini:"BOO21" cfg:"BOO21" cfgDefault:"false"`
	Boo22 bool `json:"BOO22" ini:"BOO22" cfg:"BOO22" cfgDefault:"false"`
	Boo23 bool `json:"BOO23" ini:"BOO23" cfg:"BOO23" cfgDefault:"false"`

	Boo31 bool `json:"BOO31" ini:"BOO31" cfg:"BOO31" cfgDefault:"true"`
	Boo32 bool `json:"BOO32" ini:"BOO32" cfg:"BOO32" cfgDefault:"true"`
	Boo33 bool `json:"BOO33" ini:"BOO33" cfg:"BOO33" cfgDefault:"true"`
}

type foo struct {
	Foo11 bool `json:"FOO11" ini:"FOO11" cfg:"FOO11"`
	Foo12 bool `json:"FOO12" ini:"FOO12" cfg:"FOO12"`
	Foo13 bool `json:"FOO13" ini:"FOO13" cfg:"FOO13"`

	Foo21 bool `json:"FOO21" ini:"FOO21" cfg:"FOO21" cfgDefault:"false"`
	Foo22 bool `json:"FOO22" ini:"FOO22" cfg:"FOO22" cfgDefault:"false"`
	Foo23 bool `json:"FOO23" ini:"FOO23" cfg:"FOO23" cfgDefault:"false"`

	Foo31 bool `json:"FOO31" ini:"FOO31" cfg:"FOO31" cfgDefault:"true"`
	Foo32 bool `json:"FOO32" ini:"FOO32" cfg:"FOO32" cfgDefault:"true"`
	Foo33 bool `json:"FOO33" ini:"FOO33" cfg:"FOO33" cfgDefault:"true"`

	Boo boo `json:"boo" ini:"boo" cfg:"boo"`
}

func TestBoolDefaults(t *testing.T) {
	goconfig.Path = ""
	goconfig.File = ""

	// println("\nExpected:")
	e := &foo{}
	e.Foo31 = true
	e.Foo32 = true
	e.Foo33 = true
	e.Boo.Boo31 = true
	e.Boo.Boo32 = true
	e.Boo.Boo33 = true
	// pr(e)

	// println("\nActual:")
	a := &foo{}
	goconfig.Parse(a)
	// pr(a)

	if failed(t, a, e) {
		t.Fatal()
	}
}

func TestBoolINIConfig(t *testing.T) {
	goconfig.Path = "fixtures"
	goconfig.File = "env43.ini"
	goconfig.FileRequired = true

	// println("\nExpected:")
	e := &foo{}
	e.Foo12 = true
	e.Foo22 = true
	e.Foo31 = true
	e.Foo32 = true
	e.Boo.Boo12 = true
	e.Boo.Boo22 = true
	e.Boo.Boo31 = true
	e.Boo.Boo32 = true
	// pr(e)

	// println("\nActual:")
	a := &foo{}
	goconfig.Parse(a)
	// pr(a)

	if failed(t, a, e) {
		t.Fatal()
	}
}
func TestBoolJSONConfig(t *testing.T) {
	goconfig.Path = "fixtures"
	goconfig.File = "env43.json"
	goconfig.FileRequired = true

	// println("\nExpected:")
	e := &foo{}
	e.Foo12 = true
	e.Foo22 = true
	e.Foo31 = true
	e.Foo32 = true
	e.Boo.Boo12 = true
	e.Boo.Boo22 = true
	e.Boo.Boo31 = true
	e.Boo.Boo32 = true
	// pr(e)

	// println("\nActual:")
	a := &foo{}
	goconfig.Parse(a)
	// pr(a)

	if failed(t, a, e) {
		t.Fatal()
	}
}

func failed(t *testing.T, a, e *foo) bool {
	f := false
	f = neq(t, "Foo11", a.Foo11, e.Foo11) || f
	f = neq(t, "Foo12", a.Foo12, e.Foo12) || f
	f = neq(t, "Foo13", a.Foo13, e.Foo13) || f

	f = neq(t, "Foo21", a.Foo21, e.Foo21) || f
	f = neq(t, "Foo22", a.Foo22, e.Foo22) || f
	f = neq(t, "Foo23", a.Foo23, e.Foo23) || f

	f = neq(t, "Foo31", a.Foo31, e.Foo31) || f
	f = neq(t, "Foo32", a.Foo32, e.Foo32) || f
	f = neq(t, "Foo33", a.Foo33, e.Foo33) || f

	f = neq(t, "Boo11", a.Boo.Boo11, e.Boo.Boo11) || f
	f = neq(t, "Boo12", a.Boo.Boo12, e.Boo.Boo12) || f
	f = neq(t, "Boo13", a.Boo.Boo13, e.Boo.Boo13) || f

	f = neq(t, "Boo21", a.Boo.Boo21, e.Boo.Boo21) || f
	f = neq(t, "Boo22", a.Boo.Boo22, e.Boo.Boo22) || f
	f = neq(t, "Boo23", a.Boo.Boo23, e.Boo.Boo23) || f

	f = neq(t, "Boo31", a.Boo.Boo31, e.Boo.Boo31) || f
	f = neq(t, "Boo32", a.Boo.Boo32, e.Boo.Boo32) || f
	f = neq(t, "Boo33", a.Boo.Boo33, e.Boo.Boo33) || f
	return f
}

func neq(t *testing.T, name string, a, e bool) bool {
	if a != e {
		t.Logf("%s: %v, expected=%v\n", name, a, e)
		return true
	}
	return false
}

func pr(obj interface{}) {
	j, _ := json.MarshalIndent(obj, "", "    ")
	println(string(j))
}
