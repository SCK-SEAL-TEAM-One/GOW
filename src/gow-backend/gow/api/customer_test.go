package api_test

import (
	"bytes"
	. "gow/api"
	"gow/route"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CreateCustomerHandler_Input_TN_Corporation_Should_Be_Status_201_With_TN_Corporation(t *testing.T) {
	expected, _ := ioutil.ReadFile("./customerResponse.json")
	data, _ := ioutil.ReadFile("./customerRequest.json")
	request := httptest.NewRequest("POST", "/api/v1/customers", bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()
	mockGetCustomer := mockCustomerService{}
	api := CustomerAPI{
		CustomerService: &mockGetCustomer,
	}
	testRoute := route.NewRoute(CompanyAPI{}, api, QuotationAPI{})

	testRoute.ServeHTTP(writer, request)
	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != http.StatusCreated {
		t.Errorf("expect \n'%d' \nbut got it \n'%d'", http.StatusCreated, response.StatusCode)
	}
	if string(expected) != string(actual) {
		t.Errorf("expect \n'%s' \nbut got it \n'%s'", expected, actual)
	}

}

func Test_GetAllCustomerHandler_Should_Be_CustomerInfo(t *testing.T) {
	request := httptest.NewRequest("GET", "/api/v1/customers", nil)
	expected, _ := ioutil.ReadFile("./customerResponseInfo.json")
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()
	mockGetCustomer := mockCustomerService{}
	api := CustomerAPI{
		CustomerService: &mockGetCustomer,
	}
	testRoute := route.NewRoute(CompanyAPI{}, api, QuotationAPI{})

	testRoute.ServeHTTP(writer, request)
	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if string(expected) != string(actual) {
		t.Errorf("expect \n'%s' \nbut got it \n'%s'", expected, actual)
	}

}
