package goenv

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/crgimenes/goconfig/structtag"
)

// Prefix is a string that would be placed at the beginning of the generated tags.
var Prefix string

// Usage is the function that is called when an error occurs.
var Usage func()

// Setup maps and variables
func Setup(tag string, tagDefault string) {
	Usage = DefaultUsage

	structtag.Setup()
	structtag.Prefix = Prefix
	SetTag(tag)
	SetTagDefault(tagDefault)

	structtag.ParseMap[reflect.Int64] = reflectInt
	structtag.ParseMap[reflect.Int] = reflectInt
	structtag.ParseMap[reflect.Float64] = reflectFloat
	structtag.ParseMap[reflect.String] = reflectString
	structtag.ParseMap[reflect.Bool] = reflectBool
}

// SetTag set a new tag
func SetTag(tag string) {
	structtag.Tag = tag
}

// SetTagDefault set a new TagDefault to retorn default values
func SetTagDefault(tag string) {
	structtag.TagDefault = tag
}

// Parse configuration
func Parse(config interface{}) (err error) {
	err = structtag.Parse(config, "")
	return
}

// PrintDefaultsOutput changes the default output help string
var PrintDefaultsOutput string

func getNewValue(field *reflect.StructField, value *reflect.Value, tag string, datatype string) (ret string) {

	defaultValue := field.Tag.Get(structtag.TagDefault)

	// create PrintDefaults output
	tag = strings.ToUpper(tag)
	sysvar := `$` + tag
	if runtime.GOOS == "windows" {
		sysvar = `%` + tag + `%`
	}

	if defaultValue == "" {
		PrintDefaultsOutput += ` ` + sysvar + ` ` + datatype + "\n\n"
	} else {
		printDV := " (default \"" + defaultValue + "\")"
		PrintDefaultsOutput += `  ` + sysvar + ` ` + datatype + "\n\t" + printDV + "\n"
	}

	// get value from environment variable
	ret = os.Getenv(tag)
	if ret != "" {
		return
	}

	switch datatype {
	case "bool":
		if value.Bool() {
			ret = "true"
			return
		}
	case "string":
		ret = value.String()
		if ret != "" {
			return
		}
	case "int":
		ret = strconv.FormatInt(value.Int(), 10)
		if ret != "0" {
			return
		}
	case "float64":
		f := value.Float()
		ret = strconv.FormatFloat(f, 'f', -1, 64)
		if ret != "0" {
			return
		}
	}
	// get value from default settings
	ret = defaultValue
	return
}

func reflectInt(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	newValue := getNewValue(field, value, tag, "int")
	if newValue == "" {
		return
	}
	var intNewValue int64
	intNewValue, err = strconv.ParseInt(newValue, 10, 64)
	if err != nil {
		return
	}
	value.SetInt(intNewValue)
	return
}

func reflectFloat(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	newValue := getNewValue(field, value, tag, "float64")
	if newValue == "" {
		return
	}
	var floatNewValue float64
	floatNewValue, err = strconv.ParseFloat(newValue, 64)
	if err != nil {
		return
	}
	value.SetFloat(floatNewValue)
	return
}

func reflectString(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	newValue := getNewValue(field, value, tag, "string")
	if newValue == "" {
		return
	}
	value.SetString(newValue)
	return
}

func reflectBool(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	newValue := getNewValue(field, value, tag, "bool")
	if newValue == "" {
		return
	}
	newBoolValue := newValue == "true" || newValue == "t"
	value.SetBool(newBoolValue)
	return
}

// PrintDefaults print the default help
func PrintDefaults() {
	fmt.Println("Environment variables:")
	fmt.Println(PrintDefaultsOutput)
}

// DefaultUsage is assigned for Usage function by default
func DefaultUsage() {
	fmt.Println("Usage")
	PrintDefaults()
}
