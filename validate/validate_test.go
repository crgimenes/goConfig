package validate

import (
	"testing"
	"time"
)

type testStruct struct {
	A int    `cfg:"A" cfgDefault:"100"`
	B string `cfg:"B" cfgDefault:"200"`
	C string
	D bool `cfg:"D" cfgDefault:"true"`
	F float64
	G float64 `cfg:"G" cfgDefault:"3.05"`
	H time.Duration
	I int64
	N string `cfg:"-"`
	M int
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
	B string `cfg:"SOMENAME" cfgDefault:"" cfgRequired:"true"`
}

func TestParse(t *testing.T) {

	Setup("cfg", "cfgDefault")

	s := &testStruct{A: 1, F: 1.0, S: testSub{A: 1, B: "", S: testSubSub{B: "test"}}}
	err := Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	s = &testStruct{A: 1, F: 1.0, S: testSub{A: 1, B: ""}}
	err = Parse(s)
	if err == nil {
		t.Fatal("expected error but got nil")
	}

}
