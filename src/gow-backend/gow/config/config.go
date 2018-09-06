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
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
}

func (m Mysql) GetURI() string {
	fmt.Println("GetURI PORT:", m.Port)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", m.Username, m.Password, m.Host, m.Port, m.Database)
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
	if os.Getenv("PORT") != "" {
		config.Port = os.Getenv("PORT")
	}
	if os.Getenv("MYSQL_USERNAME") != "" {
		config.Mysql.Username = os.Getenv("MYSQL_USERNAME")
	}
	if os.Getenv("MYSQL_PASSWORD") != "" {
		config.Mysql.Password = os.Getenv("MYSQL_PASSWORD")
	}
	if os.Getenv("MYSQL_DATABASE") != "" {
		config.Mysql.Database = os.Getenv("MYSQL_DATABASE")
	}
	if os.Getenv("MYSQL_HOST") != "" {
		config.Mysql.Host = os.Getenv("MYSQL_HOST")
	}
	if os.Getenv("MYSQL_PORT") != "" {
		config.Mysql.Port = os.Getenv("MYSQL_PORT")
	}
	return config, err
}
