package main

import (
	"database/sql"
	"fmt"
	"gow/api"
	configLibrary "gow/config"
	"gow/repository"
	"gow/route"
	"gow/service"
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
	companyAPI := api.CompanyAPI{
		CompanyService: &companyService,
	}
	route := route.NewRoute(companyAPI, customerAPI, api.QuotationAPI{})
	route.Run(fmt.Sprintf(":%s", config.Port))

}
