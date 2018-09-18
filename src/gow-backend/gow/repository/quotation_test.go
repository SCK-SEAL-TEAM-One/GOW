package repository_test

import (
	"database/sql"
	"gow/model"
	. "gow/repository"
	"testing"
)

func Test_InsertQuotation_Input_QuotationForm_And_Payment_And_QuotationNumber_And_VatRate_Should_Be_QuotationID_1(t *testing.T) {
	url := "root:sckshuhari@tcp(localhost:3306)/GOW?parseTime=true"
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	expectedQuotationId := 1
	quotationForm := model.QuotationForm{
		CompanyTaxID:  "0105561001221",
		CustomerTaxID: "0105537133562",
		Contact: model.Contact{
			Name:        "Nopparat Slisatkorn",
			Email:       "nopparat.slisatkorn@krungsri.com",
			PhoneNumber: "",
		},
		ProjectName: "โครงการ 3 Days Software Testing in Action Workshop",
		IncludeVAT:  false,
	}
	vatRate := 7.00
	quotationNumber := "QT256104-101002"

	quotationRepository := QuotationRepositoryMySQL{
		DBConnection: db,
	}

	quotationID, err := quotationRepository.InsertQuotation(quotationForm, quotationNumber, vatRate)
	if err != nil {
		t.Errorf("%s", err)
	}
	if int64(expectedQuotationId) != quotationID {
		t.Errorf("expect %d but it got %d", expectedQuotationId, quotationID)
	}
}

func Test_GetQuotationByQuotationNumber_Input_QuotationNumber_QT256104_101002_Should_Be_QuotationNmber_QT256104_101002(t *testing.T) {
	url := "root:sckshuhari@tcp(localhost:3306)/GOW?parseTime=true"
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	expectedQuotation := model.Quotation{
		QuotationID:     1,
		QuotationNumber: "QT256104-101002",
		CustomerTaxID:   "0105537133562",
		CompanyTaxID:    "0105561001221",
	}

	quotationNumber := "QT256104-101002"
	quotationRepository := QuotationRepositoryMySQL{
		DBConnection: db,
	}

	quotation, _ := quotationRepository.GetByQuotationNumber(quotationNumber)

	if expectedQuotation.QuotationID != quotation.QuotationID {
		t.Errorf("expect %v \nbut it got\n %v", expectedQuotation.QuotationID, quotation.QuotationID)
	}
}
func Test_GetQuotationByID_Input_QuotationID_1_Should_Be_QuotationID_1(t *testing.T) {
	url := "root:sckshuhari@tcp(localhost:3306)/GOW?parseTime=true"
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	expectedQuotation := model.Quotation{
		QuotationID:     1,
		QuotationNumber: "QT256104-101002",
	}

	quotationID := int64(1)
	quotationRepository := QuotationRepositoryMySQL{
		DBConnection: db,
	}

	quotation, _ := quotationRepository.GetQuotationByID(quotationID)

	if expectedQuotation.QuotationID != quotation.QuotationID {
		t.Errorf("expect %v \nbut it got\n %v", expectedQuotation.QuotationID, quotation.QuotationID)
	}
	if expectedQuotation.QuotationNumber != quotation.QuotationNumber {
		t.Errorf("expect %v \nbut it got\n %v", expectedQuotation.QuotationNumber, quotation.QuotationNumber)
	}
}
