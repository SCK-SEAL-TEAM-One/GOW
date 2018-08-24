package main

import (
	"database/sql"
	apiLibrary "gow-backend/api"
	"gow-backend/repository"
	"gow-backend/route"
	"gow-backend/service"
)

func main() {
	db, err := sql.Open("mysql", "root:sckshuhari@/GOW?parseTime=true")
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
	route.Run(":3000")
}
