package api_test

import (
	"bytes"
	"encoding/json"
	"gow-backend/api"
	"gow-backend/model"
	"gow-backend/route"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

type mockCustomerService struct {
}

func (mhs mockCustomerService) GetCustomers() ([]model.CustomerInfo, error) {
	return []model.CustomerInfo{}, nil
}
func (mhs mockCustomerService) CreateNewCustomer(newcustomer model.NewCustomer) (model.CustomerInfo, error) {
	var customerInfo model.CustomerInfo
	customers, _ := ioutil.ReadFile("./customerResponse.json")
	json.Unmarshal(customers, &customerInfo)
	return customerInfo, nil
}

func Test_CreateCustomerHandler_Input_TN_Corporation_Should_Be_Status_201_With_TN_Corporation(t *testing.T) {
	expected, _ := ioutil.ReadFile("./customerResponse.json")
	data, _ := ioutil.ReadFile("./customerRequest.json")
	request := httptest.NewRequest("POST", "/api/v1/customers", bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()
	mockGetCustomer := mockCustomerService{}
	api := api.API{
		CustomerService: &mockGetCustomer,
	}
	testRoute := route.NewRoute(api)
	testRoute.ServeHTTP(writer, request)
	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)
	if string(expected) != string(actual) {
		t.Errorf("expect \n'%s' \nbut got it \n'%s'", expected, actual)
	}

}
