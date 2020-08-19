package requests

type GetBankRequest struct {
	MSPId   string `json:"mspId" valid:"required"`
	Address string `json:"address" valid:"required,validHex40"`
}
