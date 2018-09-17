package api

import (
	"gow/model"
	"gow/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuotationAPI struct {
	QuotationService service.QuotationService
}

func (quotationAPI QuotationAPI) CreateQuotationHandler(context *gin.Context) {
	var quotationForm model.QuotationForm
	err := context.Bind(&quotationForm)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	quotationInfo, err := quotationAPI.QuotationService.CreateQuotation(quotationForm)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, quotationInfo)
}
