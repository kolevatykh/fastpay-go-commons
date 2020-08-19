package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/invoice-state-enum"

type UpdateInvoiceRequest struct {
	Number       string                          `json:"number" valid:"required,lte(255)"`
	Recipient    string                          `json:"recipient" valid:"required,validHex40"`
	Payer        string                          `json:"payer" valid:"required,validHex40"`
	State        invoice_state_enum.InvoiceState `json:"state" valid:"required"`
	CurrencyCode int                             `json:"currencyCode" valid:"required,gte(0),lte(999)"`
	ErrorCode    int                             `json:"errorCode"`
}
