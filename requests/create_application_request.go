package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
	"github.com/SolarLabRU/fastpay-go-commons/requests"
)

type CreateApplicationRequest struct {
	Id      string           `json:"id" validate:"required,uuid"`
	Owner   string           `json:"owner" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Terms   models.TermsDeal `json:"terms" validate:"required"`
	MsgHash string           `json:"msgHash" validate:"required"`
	Sig     requests.SignDto `json:"sig" validate:"required"`
	Exp     int64            `json:"exp" valid:"required~ErrorTimestampNotPassed"`
}
