package repository

import (
	"database/sql"
	"fmt"
	"gow/model"
	"time"
)

type QuotationRepository interface {
	InsertQuotation(model.QuotationForm, model.Payment, float64) (int64, error)
	GetQuotationByID(int64) (model.Quotation, error)
}

type QuotationRepositoryMySQL struct {
	DBConnection *sql.DB
}

func (quotationRepository QuotationRepositoryMySQL) GenerateQuotationNumber() string {
	var quatationNumber int
	statementQuery := `SELECT COUNT(*) FROM quotation WHERE YEAR(created_time) = ?`
	row := quotationRepository.DBConnection.QueryRow(statementQuery, time.Now().Year())
	row.Scan(&quatationNumber)
	return fmt.Sprintf("QT%d%d-%06d", time.Now().Year(), time.Now().Month(), quatationNumber+1)
}

func (quotationRepository QuotationRepositoryMySQL) InsertQuotation(quotationForm model.QuotationForm, payment model.Payment, vatRate float64) (int64, error) {
	newQuotation := model.Quotation{
		QuotationNumber:    quotationRepository.GenerateQuotationNumber(),
		CompanyID:          quotationForm.CompanyID,
		CustomerID:         quotationForm.CustomerID,
		ContactName:        quotationForm.Contact.Name,
		ContactEmail:       quotationForm.Contact.Email,
		ContactPhoneNumber: quotationForm.Contact.PhoneNumber,
		TotalPrice:         payment.TotalPrice,
		Discount:           payment.Discount,
		PriceAfterDiscount: payment.PriceAfterDiscount,
		VAT:                payment.VAT,
		NetTotalPrice:      payment.NetTotalPrice,
		TotalPriceThai:     payment.TotalPriceThai,
		ProjectName:        quotationForm.ProjectName,
		QuotationDate:      time.Now(),
		VatRate:            vatRate,
		VatIncluded:        quotationForm.IncludeVAT,
	}

	statement := `INSERT INTO quotation 
	(quotation_number,customer_id,company_id,contact_name,contact_email,contact_phoneNumber,
	total_price,discount_price,price_after_discount,vat,net_total_price,total_price_thai,
	project_name,quotation_date,vat_rate,vat_included,created_time,updated_time)
	VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`
	statementInsert, err := quotationRepository.DBConnection.Prepare(statement)
	if err != nil {
		return 0, err
	}
	defer statementInsert.Close()
	result, err := statementInsert.Exec(
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

func (quotationRepository QuotationRepositoryMySQL) GetQuotationByID(quotationID int64) (model.Quotation, error) {
	var quotation model.Quotation
	statementQuery := `SELECT (id, quotation_number, customer_id, company_id, contact_name, contact_email, contact_phoneNumber,
		total_price,discount_price,price_after_discount,vat,net_total_price,total_price_thai,
		project_name,quotation_date,vat_rate,vat_included,created_time,updated_time) FROM quotation WHERE id = ?`
	row := quotationRepository.DBConnection.QueryRow(statementQuery, quotationID)
	err := row.Scan(
		&quotation.QuotationID,
		&quotation.QuotationNumber,
		&quotation.CustomerID,
		&quotation.CompanyID,
		&quotation.ContactName,
		&quotation.ContactEmail,
		&quotation.ContactPhoneNumber,
		&quotation.TotalPrice,
		&quotation.Discount,
		&quotation.PriceAfterDiscount,
		&quotation.VAT,
		&quotation.NetTotalPrice,
		&quotation.TotalPriceThai,
		&quotation.ProjectName,
		&quotation.QuotationDate,
		&quotation.VatRate,
		&quotation.VatIncluded,
		&quotation.CreatedTime,
		&quotation.UpdatedTime,
	)
	return quotation, err
}
