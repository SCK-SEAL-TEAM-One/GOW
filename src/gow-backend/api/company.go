package api

import (
	"gow-backend/model"
	"gow-backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyAPI struct {
	CompanyService service.CompanyService
}

func (api CompanyAPI) CreateCompanyHandler(c *gin.Context) {
	var newCompany model.NewCompany
	err := c.Bind(&newCompany)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	createdCompany, err := api.CompanyService.CreateNewCompany(newCompany)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, createdCompany)
}
