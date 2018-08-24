package api

import (
	"gow-backend/model"
	"gow-backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	CustomerService service.CustomerService
}

func (api API) CreateCustomerHandler(c *gin.Context) {
	var newCustomer model.NewCustomer
	err := c.Bind(&newCustomer)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	customers, err := api.CustomerService.CreateNewCustomer(newCustomer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, customers)
}

func (api API) GetAllCustomerHandler(c *gin.Context) {
	var customer model.CustomerInfo
	err := c.Bind(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	customers, err := api.CustomerService.ListCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, customers)

}
