package route

import (
	apiLibrary "gow/api"

	"github.com/gin-gonic/gin"
)

func NewRoute(companyApi apiLibrary.CompanyAPI, customerApi apiLibrary.CustomerAPI) *gin.Engine {
	route := gin.Default()
	route.POST("api/v1/customers", customerApi.CreateCustomerHandler)
	route.GET("api/v1/customers", customerApi.GetAllCustomerHandler)
	route.POST("api/v1/companies", companyApi.CreateCompanyHandler)
	return route
}
