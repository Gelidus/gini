package gini

import (
	"os"
	"reflect"
	"testing"
)

type Config struct {
	Section struct {
		Ahoy   string
		Number int32
	}

	Application struct {
		GridSize int32
		Formal   string
	}
}

func TestRead(t *testing.T) {
	conf := &Config{}

	file, err := os.Open("test.ini")
	if err != nil {
		t.Fatal(err)
	}

	err = Read(conf, file)
	if err != nil {
		t.Fatal(err)
	}

	expected := &Config{
		Section: struct {
			Ahoy   string
			Number int32
		}{
			Ahoy:   "hello",
			Number: 42,
		},
		Application: struct {
			GridSize int32
			Formal   string
		}{
			GridSize: 100,
			Formal:   "hello world",
		},
	}

	ok := reflect.DeepEqual(conf, expected)
	if !ok {
		t.Errorf("expected %v, got %v", expected, conf)
	}
}

func TestReadFile(t *testing.T) {
	conf := &Config{}

	err := ReadFile(conf, "test.ini")
	if err != nil {
		t.Fatal(err)
	}

	expected := &Config{
		Section: struct {
			Ahoy   string
			Number int32
		}{
			Ahoy:   "hello",
			Number: 42,
		},
		Application: struct {
			GridSize int32
			Formal   string
		}{
			GridSize: 100,
			Formal:   "hello world",
		},
	}

	ok := reflect.DeepEqual(conf, expected)
	if !ok {
		t.Errorf("expected %v, got %v", expected, conf)
	}
}
