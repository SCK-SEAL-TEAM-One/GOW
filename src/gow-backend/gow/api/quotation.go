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
	quotationID, err := quotationAPI.QuotationService.CreateQuotation(quotationForm)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	quotationInfo, err := quotationAPI.QuotationService.GetQuotationByQuotationID(quotationID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, quotationInfo)
}

func (quotationAPI QuotationAPI) GetQuotationHandler(context *gin.Context) {
}
