package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Mysql string `json:"mysql"`
	Port  string `json:"port"`
}

func SetupConfig() (Config, error) {
	var config Config
	environment := "development"
	if os.Getenv("ENV") != "" {
		environment = os.Getenv("ENV")
	}

	configFile, err := ioutil.ReadFile(fmt.Sprintf("./configs/%s.json", environment))
	if err != nil {
		fmt.Printf("cannot read config file %s", err)
		return config, err
	}

	err = json.Unmarshal(configFile, &config)
	return config, err
}
