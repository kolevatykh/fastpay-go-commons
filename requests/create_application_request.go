package requests

type CreateApplicationRequest struct {
	Id                       string  `json:"id" validate:"required,uuid"`
	AddressInitiator         string  `json:"addressInitiator" validate:"required,validHex40"`
	AddressAcceptor          string  `json:"addressAcceptor" validate:"required,validHex40"`
	AmountVotesCloseContract int     `json:"amountVotesCloseContract" validate:"required,min=0"`
	MsgHash                  string  `json:"msgHash" validate:"required"`
	Sig                      SignDto `json:"sig" validate:"required,dive"`
	Exp                      int64   `json:"exp" validate:"required"`
}
