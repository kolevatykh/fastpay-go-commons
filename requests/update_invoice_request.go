package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/invoice-state-enum"

type UpdateInvoiceRequest struct {
	Number       string                          `json:"number" validate:"required,lte=255"`
	Recipient    string                          `json:"recipient" validate:"required,validHex40"`
	Payer        string                          `json:"payer" validate:"required,validHex40"`
	State        invoice_state_enum.InvoiceState `json:"state" validate:"required"`
	CurrencyCode int                             `json:"currencyCode" validate:"required,gte=0,lte=999"`
	ErrorCode    int                             `json:"errorCode"`
}
