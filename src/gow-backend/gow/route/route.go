package route

import (
	apiLibrary "gow/api"

	"github.com/gin-gonic/gin"
)

func NewRoute(api apiLibrary.API) *gin.Engine {
	route := gin.Default()
	route.POST("api/v1/customers", api.CreateCustomerHandler)
	route.GET("api/v1/customers", api.GetAllCustomerHandler)
	return route
}
