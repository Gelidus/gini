package gini

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/vaughan0/go-ini"
)

// Read will load configuration specified by reader
// returning reflected structure of given configuration
func Read(config interface{}, r io.Reader) error {
	file, err := ini.Load(r)
	if err != nil {
		return err
	}

	return reflectConfig(config, &file)
}

// ReadFile will load configuration from given file name
// returning reflected structure of given configuration
func ReadFile(config interface{}, fileName string) error {
	file, err := ini.LoadFile(fileName)
	if err != nil {
		return err
	}

	return reflectConfig(config, &file)
}

// reflectConfig reflects given ini.File into the given interface
func reflectConfig(config interface{}, file *ini.File) error {
	root := reflect.ValueOf(config).Elem()

	// for each section
	for i := 0; i < root.NumField(); i++ {
		section := root.Field(i)
		sectionName := strings.ToLower(root.Type().Field(i).Name)

		// for each field in section
		for y := 0; y < section.NumField(); y++ {
			field := section.Field(y)
			name := strings.ToLower(section.Type().Field(y).Name)

			// retrieve from file map
			value, ok := file.Get(sectionName, name)
			if !ok {
				fmt.Println("Variable is missing from config file:", sectionName, "->", name)
				continue
			}

			// set field by kind
			switch field.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				num, err := strconv.ParseInt(value, 10, field.Type().Bits())
				if err != nil {
					return err
				}
				field.SetInt(num)

			case reflect.String:
				field.SetString(value)

			case reflect.Bool:
				b, err := strconv.ParseBool(value)
				if err != nil {
					return err
				}
				field.SetBool(b)
			}
		}
	}

	return nil
}
