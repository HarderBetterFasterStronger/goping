package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// Links urls of websites to check
type Links []string

// Config JSON block
type Config struct {
	Links Links `json:"links"`
}

// ConfigLoader interface for loading configuration file
type ConfigLoader interface {
	load() (Config, error)
}

func (c Config) load() (Config, error) {
	// read args
	configFileName := flag.String("c", "config.json", "json configuration file.  Object with with 'links' array")
	flag.Parse()
	if configFileName == nil {
		fmt.Println("config file not found or empty")
		os.Exit(1)
	}

	// Open our jsonFile
	jsonFile, err := os.Open(*configFileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Successfully Opened: " + *configFileName)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Could not read contents of file: " + *configFileName)
		os.Exit(1)
	}

	err = json.Unmarshal([]byte(byteValue), &c)
	if err != nil {
		fmt.Println("JSON did not parse configuration file: " + *configFileName)
		os.Exit(1)
	}

	return c, err

}
