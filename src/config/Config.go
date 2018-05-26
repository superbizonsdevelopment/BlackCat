package config

import(
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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
	
	configFile, err := ioutil.ReadFile("./config.json")

	
	if err != nil {
		log.Println(err.Error())
		
		configFile, err := os.Create("./config.json")
		
		if err != nil {
			log.Println(err.Error())
			return err
		}
		
		defer configFile.Close()
		
		return err
	}
	
	log.Println(string(configFile))
	
	err = json.Unmarshal(configFile, &config)
	
	if err != nil {
		log.Println(err.Error())
		return err
	}
	
	Token = config.Token
	Prefix = config.Prefix
	
	return nil
}