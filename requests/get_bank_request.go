package requests

type GetBankRequest struct {
	MSPId   string `json:"mspId" valid:"required"`
	Address string `json:"address" valid:"required~60364,validHex40~60367"`
}
