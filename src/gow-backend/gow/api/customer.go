package api

import (
	"gow/model"
	"gow/service"
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
	createdCustomer, err := api.CustomerService.CreateNewCustomer(newCustomer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, createdCustomer)
}

func (api API) GetAllCustomerHandler(c *gin.Context) {
	customers, err := api.CustomerService.ListCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, customers)

}
