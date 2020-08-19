package requests

type WithdrawRequest struct {
	AddressFrom   string  `json:"addressFrom" valid:"required~60302,validHex40~60301"`
	Amount        int64   `json:"amount" valid:"required~60313"`
	CurrencyCode  int     `json:"currencyCode" valid:"required~60318,range(0|999)~60317"`
	Payload       string  `json:"payload"`
	TransactionId string  `json:"transactionId" valid:"required~60331,uuidv4"`
	MsgHash       string  `json:"msgHash" valid:"required"`
	Exp           int64   `json:"exp" valid:"required~60332"`
	Sig           SignDto `json:"sig" valid:"required"`
}
