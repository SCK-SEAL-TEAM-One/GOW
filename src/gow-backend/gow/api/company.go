package api

import (
	"gow/model"
	"gow/service"
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

func (api CompanyAPI) GetAllCompaniesHandler(c *gin.Context) {
	companies, err := api.CompanyService.GetCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, companies)
}
