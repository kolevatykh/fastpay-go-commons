package requests

type CreateInvoiceRequest struct {
	Number       string `json:"number" valid:"required~60352,maxstringlength(255)~60353"`
	CurrencyCode int    `json:"currencyCode" valid:"required~60318,range(0|999)~60317"`
	Amount       int64  `json:"amount" valid:"required~60313,range(1|9223372036854775807)"`
	Description  string `json:"description"`
	Recipient    string `json:"recipient" valid:"required~60302,validHex40~60301"`
	Payer        string `json:"payer" valid:"validHex40~60301"`
}
