package requests

type TopupRequest struct {
	AddressTo     string `json:"addressTo" validate:"required,validHex40"`
	Amount        int64  `json:"amount" validate:"required"`
	CurrencyCode  int    `json:"currencyCode" validate:"required,min=0"`
	Payload       string `json:"payload"`
	TransactionId int    `json:"transactionId" validate:"required,min=0"`
}
