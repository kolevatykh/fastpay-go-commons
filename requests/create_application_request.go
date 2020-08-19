package requests

type CreateApplicationRequest struct {
	Id                       string  `json:"id" valid:"required,uuid"`
	AddressInitiator         string  `json:"addressInitiator" valid:"required,validHex40"`
	AddressAcceptor          string  `json:"addressAcceptor" valid:"required,validHex40"`
	AmountVotesCloseContract int     `json:"amountVotesCloseContract" valid:"required,min(0)"`
	MsgHash                  string  `json:"msgHash" valid:"required"`
	Sig                      SignDto `json:"sig" valid:"required"`
	Exp                      int64   `json:"exp" valid:"required"`
}
