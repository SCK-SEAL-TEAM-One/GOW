package main

import (
	"database/sql"
	"fmt"
	"gow-backend/api"
	configLibrary "gow-backend/config"
	"gow-backend/repository"
	"gow-backend/route"
	"gow-backend/service"
)

func main() {
	config, err := configLibrary.SetupConfig()
	if err != nil {
		panic(err.Error())
	}

	db, err := sql.Open("mysql", config.Mysql.GetURI())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	customerRepository := repository.CustomerRepositoryMySQL{
		DBConnection: db,
	}
	customerService := service.CustomerServiceMySQL{
		CustomerRepository: &customerRepository,
	}
	customerAPI := api.CustomerAPI{
		CustomerService: &customerService,
	}
	route := route.NewRoute(api.CompanyAPI{}, customerAPI)
	route.Run(fmt.Sprintf(":%s", config.Port))
}
