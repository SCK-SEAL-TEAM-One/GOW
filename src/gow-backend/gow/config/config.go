package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Mysql Mysql  `json:"mysql" yaml:"mysql"`
	Port  string `json:"port" yaml:"port"`
}
type Mysql struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
}

func (m Mysql) GetURI() string {
	return fmt.Sprintf("%s:%s@/%s?parseTime=true", m.Username, m.Password, m.Database)
}

func SetupConfig() (Config, error) {
	var config Config
	environment := "development"
	if os.Getenv("ENV") != "" {
		environment = os.Getenv("ENV")
	}

	configFile, err := ioutil.ReadFile(fmt.Sprintf("../configs/%s.yml", environment))
	if err != nil {
		fmt.Printf("cannot read config file %s", err)
		return config, err
	}
	err = yaml.Unmarshal(configFile, &config)
	return config, err
}
