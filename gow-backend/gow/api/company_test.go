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

func Test_CreateCompanyHandler_Input_Siamchamnankit_Should_Be_Status_201_And_Company_Siamchamnankit(t *testing.T) {
	expected, _ := ioutil.ReadFile("companyResponse.json")
	data, _ := ioutil.ReadFile("companyRequest.json")
	request := httptest.NewRequest("POST", "/api/v1/companies", bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()

	companyService := mockCompanyService{}
	CompanyAPI := CompanyAPI{
		CompanyService: &companyService,
	}
	testRoute := route.NewRoute(CompanyAPI, CustomerAPI{})

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

func Test_GetAllCompaniesHandler_Should_Be_CompanyInfo(t *testing.T) {
	expected, _ := ioutil.ReadFile("companiesResponseInfo.json")

	request := httptest.NewRequest("GET", "/api/v1/companies", nil)
	writer := httptest.NewRecorder()
	mockGetcompanies := mockCompanyService{}
	api := CompanyAPI{
		CompanyService: &mockGetcompanies,
	}
	testRoute := route.NewRoute(api, CustomerAPI{})

	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if string(expected) != string(actual) {
		t.Errorf("expect \n'%s' \nbut got it \n'%s'", expected, actual)
	}
}
