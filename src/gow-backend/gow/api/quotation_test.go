package api_test

import (
	"bytes"
	. "gow/api"
	"gow/route"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_CreateQuotationHandler_Input_RequestQuotationForm_Should_Be_Status_201_And_QuotationInfo(t *testing.T) {
	expected, _ := ioutil.ReadFile("quotationInfo.json")
	requestQuotationFormData, _ := ioutil.ReadFile("quotationForm.json")
	request := httptest.NewRequest("POST", "/api/v1/quotation", bytes.NewBuffer(requestQuotationFormData))
	writer := httptest.NewRecorder()
	quotationAPI := QuotationAPI{
		QuotationService: &mockQuotationService{},
	}
	testRoute := route.NewRoute(CompanyAPI{}, CustomerAPI{}, quotationAPI)

	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if string(expected) != string(actual) {
		t.Errorf("expect %s \nbut it got\n %s", expected, actual)
	}
}
