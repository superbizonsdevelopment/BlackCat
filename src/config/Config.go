package config

import(
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	Token string
	Prefix string
	
	config *Configuration
)

type Configuration struct {
	Token string `json:"Token"`
	Prefix string `json:"Prefix"`
}

func LoadConfig() error {
	log.Println("Reading from config file...")
	
	file, err := ioutil.ReadFile("./config.json")
	
	if err != nil {
		log.Println(err.Error())
		return err
	}
	
	log.Println(string(file))
	
	err = json.Unmarshal(file, &config)
	
	if err != nil {
		log.Println(err.Error())
		return err
	}
	
	Token = config.Token
	Prefix = config.Prefix
	
	return nil
}