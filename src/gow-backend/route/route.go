package route

import (
	apiLibrary "gow-backend/api"

	"github.com/gin-gonic/gin"
)

func NewRoute(api apiLibrary.API) *gin.Engine {
	route := gin.Default()
	route.POST("api/v1/customers", api.CreateCustomerHandler)
	return route
}
