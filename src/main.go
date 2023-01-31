package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Dashboard string `yaml:"dashboard"`
}

func main() {
	// Read the file
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a struct to hold the YAML data
	var server Server

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &server)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the data
	fmt.Println(server.Dashboard)
}
