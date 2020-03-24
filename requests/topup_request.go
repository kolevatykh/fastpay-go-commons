package requests

type TopupRequest struct {
	AddressTo     string `json:"addressTo" validate:"required,validHex40"` // TODO переиминовать в address?
	Amount        int    `json:"amount" validate:"required"`
	CurrencyCode  int    `json:"currencyCode" validate:"required,min=0"`
	Payload       string `json:"payload"`
	Timestamp     int    `json:"timestamp" validate:"required"`
	TransactionId int    `json:"transactionId" validate:"required,min=0"`
}
