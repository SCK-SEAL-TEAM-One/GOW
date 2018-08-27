package api

import (
	"gow-backend/model"
	"gow-backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerAPI struct {
	CustomerService service.CustomerService
}

func (api CustomerAPI) CreateCustomerHandler(c *gin.Context) {
	var newCustomer model.NewCustomer
	err := c.Bind(&newCustomer)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	createdCustomer, err := api.CustomerService.CreateNewCustomer(newCustomer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, createdCustomer)
}

func (api CustomerAPI) GetAllCustomerHandler(c *gin.Context) {
	customers, err := api.CustomerService.ListCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, customers)

}
