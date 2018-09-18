package repository

import (
	"database/sql"
	"gow/model"
	"time"
)

type QuotationRepository interface {
	InsertQuotation(model.QuotationForm, model.Payment, string, float64) (int64, error)
}
type QuotationRepositoryMySQL struct {
	DBConnection *sql.DB
}

func (quotationRepository QuotationRepositoryMySQL) InsertQuotation(quotationForm model.QuotationForm, payment model.Payment, quotationNumber string, vatRate float64) (int64, error) {
	newQuotation := model.Quotation{
		QuotationNumber:    quotationNumber,
		CompanyTaxID:       quotationForm.CompanyTaxID,
		CustomerTaxID:      quotationForm.CustomerTaxID,
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
	(quotation_number,customer_taxid, company_taxid,contact_name,contact_email,contact_phoneNumber,
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
		newQuotation.CustomerTaxID,
		newQuotation.CompanyTaxID,
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
func (quotationRepository QuotationRepositoryMySQL) GetByQuotationNumber(quotationNumber string) (model.Quotation, error) {
	var quotation model.Quotation
	statementQuery := `SELECT quotation_id, quotation_number, customer_taxid, company_taxid, contact_name, contact_email, contact_phoneNumber,
		total_price,discount_price,price_after_discount,vat,net_total_price,total_price_thai,
		project_name,quotation_date,vat_rate,vat_included,created_time,updated_time FROM quotation WHERE quotation_number = ?`
	row := quotationRepository.DBConnection.QueryRow(statementQuery, quotationNumber)
	err := row.Scan(
		&quotation.QuotationID,
		&quotation.QuotationNumber,
		&quotation.CustomerTaxID,
		&quotation.CompanyTaxID,
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
