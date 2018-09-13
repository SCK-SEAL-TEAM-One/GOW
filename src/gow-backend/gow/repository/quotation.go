package repository

import (
	"database/sql"
	"fmt"
	"gow/model"
	"time"
)

type QuotationRepository interface {
	InsertQuotation(model.QuotationForm, model.Payment) (int, error)
	GetQuotationByID(int) (model.Quotation, error)
}

type QuotationRepositoryMySQL struct {
	DBConnection *sql.DB
}

func (quotationRepository QuotationRepositoryMySQL) GenerateQuotationNumber() string {
	var quatationNumber int
	statementQuery := `SELECT COUNT(*) FROM quotation WHERE YEAR(created_time) = ?`
	row := comrepo.DBConnection.QueryRow(statementQuery, time.Now().Year())
	row.Scan(&quatationNumber)
	return fmt.Sprintf("QT%d%d-%06d", time.Now().Year(), time.Now().Month(), quatationNumber+1)
}

func (quotationRepository QuotationRepositoryMySQL) InsertQuotation(quotationForm model.QuotationForm, payment model.Payment, vatRate float64) (int, error) {
	newQuotation := model.Quotation{
		QuotationNumber:    quotationRepository.GenerateQuotationNumber(),
		CustomerID:         quotationForm.CompanyID,
		CustomerID:         quotationForm.CustomerID,
		ContactName:        quotationForm.ContactName,
		ContactEmail:       quotationForm.ContactEmail,
		ContactPhoneNumber: quotationForm.ContactPhoneNumber,
		TotalPrice:         payment.TotalPrice,
		Discount:           payment.Discount,
		PriceAfterDiscount: payment.PriceAfterDiscount,
		VAT:                payment.VAT,
		NetTotalPrice:      payment.NetTotalPrice,
		TotalPriceThai:     payment.TotalPriceThai,
		ProjectName:        quotationForm.ProjectName,
		QuotationDate:      time.Now(),
		VatRate:            vatRate,
		VatIncluded:        quotationForm.VatIncluded,
	}

	statement := `INSERT INTO quotation 
	(quotation_number,customer_id,company_id,contact_name,contact_email,contact_phoneNumber,
	total_price,discount_price,price_after_discount,vat,net_total_price,total_price_thai,
	project_name,quotation_date,vat_rate,vat_included,,created_time,updated_time)
	VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?, ? )`
	statementInsert, err := quotationRepository.DBConnection.Prepare(statement)
	if err != nil {
		return false, err
	}
	defer statementInsert.Close()
	result, err = statementInsert.Exec(
		newQuotation.QuotationNumber,
		newQuotation.CustomerID,
		newQuotation.CompanyID,
		newQuotation.ContactName,
		newQuotation.ContactEmail,
		newQuotation.ContactPhoneNumber,
		newQuotation.TotalPrice,
		newQuotation.Discount,
		newQuotation.PriceAfterDiscount,
		newQuotation.VAT,
		newQuotation.NetTotalPrice,
		newQuotation.TotalPriceThai,
		newQuotation.ProjectName,
		newQuotation.QuotationDate,
		newQuotation.VatRate,
		newQuotation.VatIncluded,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (quotationRepository QuotationRepositoryMySQL) GetQuotationByID(int) (model.Quotation, error) {
	return model.Quotation{}, nil
}
