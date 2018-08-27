package main

import (
	"database/sql"
	"fmt"
	apiLibrary "gow-backend/api"
	configLibrary "gow-backend/config"
	"gow-backend/repository"
	"gow-backend/route"
	"gow-backend/service"
)

func main() {
	config, err := configLibrary.SetupConfig()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	db, err := sql.Open("mysql", config.Mysql.GetURI())
	if err != nil {
		fmt.Printf("%s", err)
		return
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
