package requests

type WithdrawRequest struct {
	AddressFrom   string  `json:"addressFrom" validate:"required,validHex40"`
	Amount        int64   `json:"amount" validate:"required"`
	CurrencyCode  int     `json:"currencyCode" validate:"required,min=0"`
	Payload       string  `json:"payload"`
	TransactionId string  `json:"transactionId" validate:"required,uuid4"`
	MsgHash       string  `json:"msgHash" validate:"required"`
	Exp           int64   `json:"exp" validate:"required"`
	Sig           SignDto `json:"sig" validate:"required,dive"`
}
