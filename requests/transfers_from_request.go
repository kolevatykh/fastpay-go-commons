package requests

type TransfersFromRequest struct {
	AddressFrom   string            `json:"addressFrom" validate:"required,validHex40"`
	To            []TransfersFromTo `json:"to" validate:"required,dive"`
	CurrencyCode  int               `json:"currencyCode" validate:"required,min=0"`
	TransactionId string            `json:"transactionId" validate:"required,uuid4"`
}

type TransfersFromTo struct {
	To      string `json:"to" validate:"required,validHex40or64"`
	Amount  int64  `json:"amount" validate:"required"`
	Payload string `json:"payload"`
}
