package responses

type BaseResponse struct {
	ErrorCode     int    `json:"errorCode"`
	ErrorMessage  string `json:"errorMessage"`
	TransactionId string `json:"transactionId"`
	SenderAddress string `json:"senderAddress"`
}
