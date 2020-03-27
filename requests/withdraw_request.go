package requests

type WithdrawRequest struct {
	AddressFrom   string  `json:"addressFrom" validate:"required,validHex40"`
	Amount        int64   `json:"amount" validate:"required"`
	CurrencyCode  int     `json:"currencyCode" validate:"required,min=0"`
	Payload       string  `json:"payload"`
	TransactionId int     `json:"transactionId" validate:"required,min=0"`
	MsgHash       string  `json:"msgHash" validate:"required"`
	Sig           SignDto `json:"sig" validate:"required,dive"`
}
