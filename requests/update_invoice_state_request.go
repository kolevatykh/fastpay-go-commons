package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/invoice-state-enum"

type UpdateInvoiceStateRequest struct {
	Number       string                          `json:"number" valid:"required,lte(255)"`
	Recipient    string                          `json:"recipient" valid:"required,validHex40"`
	CurrencyCode int                             `json:"currencyCode" valid:"required,gte(0),lte(999)"`
	State        invoice_state_enum.InvoiceState `json:"state" valid:"required"`
}
