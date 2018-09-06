package route

import (
	"fmt"
	apiLibrary "gow/api"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRoute(companyApi apiLibrary.CompanyAPI, customerApi apiLibrary.CustomerAPI) *gin.Engine {
	route := gin.Default()
	route.POST("api/v1/customers", customerApi.CreateCustomerHandler)
	route.GET("api/v1/customers", customerApi.GetAllCustomerHandler)
	route.POST("api/v1/companies", companyApi.CreateCompanyHandler)
	route.GET("api/v1/companies", companyApi.GetAllCompaniesHandler)
	route.GET("/health", func(c *gin.Context) {

		startTime := time.Now()

		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		megaByte := TransformByteToMegabtye(m.Sys)

		t := time.Since(startTime)
		c.JSON(http.StatusOK, gin.H{
			"status":   "OK",
			"duration": fmt.Sprintf("%.2f ms", float64(t)/1000),
			"memory":   fmt.Sprintf("%d MB", megaByte),
		})
	})
	return route
}

func TransformByteToMegabtye(b uint64) uint64 {
	return (b / 1024) / 1024

}
