package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Dashboard string `yaml:"dashboard"`
}

func sendStatus(myurl string, host string, status string) {

	values := map[string]string{"host": host, "status": status}
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(myurl, "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
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

	// Get the hostname running the agent.
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}
	url := server.Dashboard

	sendStatus(url, hostname, "pass")
}
