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
	companyRepository := repository.CompanyRepositoryMySQL{
		DBConnection: db,
	}
	quotationRepository := repository.QuotationRepositoryMySQL{
		DBConnection: db,
	}
	orderRepository := repository.OrderRepositoryMySQL{
		DBConnection: db,
	}

	customerService := service.CustomerServiceMySQL{
		CustomerRepository: &customerRepository,
	}
	companyService := service.CompanyServiceMySQL{
		CompanyRepository: &companyRepository,
	}

	quotationService := service.QuotationServiceMySQL{
		QuotationRepository: &quotationRepository,
		OrderRepository:     &orderRepository,
		CompanyService:      &companyService,
		CustomerService:     &customerService,
	}

	customerAPI := api.CustomerAPI{
		CustomerService: &customerService,
	}
	companyAPI := api.CompanyAPI{
		CompanyService: &companyService,
	}
	quotationAPI := api.QuotationAPI{
		QuotationService: &quotationService,
	}
	route := route.NewRoute(companyAPI, customerAPI, quotationAPI)
	route.Run(fmt.Sprintf(":%s", config.Port))

}
