package requests

type SetPaymentErrorRequest struct {
	Number       string `json:"number" valid:"required~60352,maxstringlength(255)~60353"`
	Recipient    string `json:"recipient" valid:"required~60302,validHex40~60301"`
	CurrencyCode int    `json:"currencyCode" valid:"required~60318,range(0|999)~60317"`
	ErrorCode    int    `json:"errorCode" valid:"required"`
}
