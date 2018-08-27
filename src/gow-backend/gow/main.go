package main

import (
	"database/sql"
	"fmt"
	"gow-backend/gow/api"
	"gow-backend/gow/route"
	"gow-backend/gow/service"
	configLibrary "gow/config"
	"gow/repository"
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

	companyRepository := repository.CompanyRepositoryMySQL{
		DBConnection: db,
	}
	companyService := service.CompanyServiceMySQL{
		CompanyRepository: &companyRepository,
	}

	route := route.NewRoute(api.CompanyAPI{}, customerAPI)
	route.Run(fmt.Sprintf(":%s", config.Port))

}
