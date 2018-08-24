package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	apiLibrary "gow-backend/api"
	"gow-backend/repository"
	"gow-backend/route"
	"gow-backend/service"
	"io/ioutil"
	"os"
)

type Config struct {
	Mysql string `json:"mysql"`
	Port  string `json:"port"`
}

func main() {
	var config Config
	environment := "development"
	if os.Getenv("ENV") != "" {
		environment = os.Getenv("ENV")
	}

	configFile, err := ioutil.ReadFile(fmt.Sprintf("./configs/%s.json", environment))
	if err != nil {
		fmt.Printf("cannot read config file %s", err)
		return
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Printf("cannot Unmarshal config %s", err)
		return
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s?parseTime=true", config.Mysql))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	customerRepository := repository.CustomerRepositoryMySQL{
		ConnetionDB: db,
	}
	customerService := service.CustomerServiceMySQL{
		CustomerRepository: &customerRepository,
	}
	api := apiLibrary.API{
		CustomerService: &customerService,
	}
	route := route.NewRoute(api)
	route.Run(fmt.Sprintf(":%s", config.Port))
}
