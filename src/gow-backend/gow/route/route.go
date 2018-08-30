package route

import (
	apiLibrary "gow/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoute(companyApi apiLibrary.CompanyAPI, customerApi apiLibrary.CustomerAPI) *gin.Engine {
	route := gin.Default()
	route.POST("api/v1/customers", customerApi.CreateCustomerHandler)
	route.GET("api/v1/customers", customerApi.GetAllCustomerHandler)
	route.POST("api/v1/companies", companyApi.CreateCompanyHandler)
	route.GET("api/v1/companies", companyApi.GetAllCompaniesHandler)
	route.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
	return route
}
