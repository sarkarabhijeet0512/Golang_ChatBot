package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config is the struct for yaml files
type Config struct {
	VerifyToken string `yaml:"verify_token"`
	AccessToken string `yaml:"access_token"`
	AppSecret   string `yaml:"app_secret"`
}

func parseContentFile() string {
	contentFile, err := ioutil.ReadFile("content.yml")
	if err != nil {
		log.Println("Error opening content file", err)
	}
	er, _ := yaml.Marshal(contentFile)
	if er != nil {
		log.Println("Couldnt Marshall file", er)

	}
	return string(er)
}
func (c *Config) readYaml() {
	yamlFile, err := ioutil.ReadFile("bot.config.yml")
	if err != nil {
		log.Println("Error reading config", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Println("Couldnt Marshall file", err)
	}
	// return c
}
func getToken() string {
	var c Config
	c.readYaml()
	v, err := json.Marshal(c)
	if err != nil {
		log.Println("Couldn't Marshalling our json file", err)
	}
	return string(v)
}
