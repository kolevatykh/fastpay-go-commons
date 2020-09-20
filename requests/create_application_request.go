package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
)

type CreateApplicationRequest struct {
	Id      string           `json:"id" valid:"required,uuid"`
	OfferId string           `json:"offerId"`
	Owner   string           `json:"owner" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Terms   models.TermsDeal `json:"terms" valid:"required"`
	MsgHash string           `json:"msgHash" valid:"required"`
	Sig     SignDto          `json:"sig" valid:"required"`
	Exp     int64            `json:"exp" valid:"required~ErrorTimestampNotPassed"`
}
