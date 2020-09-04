package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/member-deal-type-enum"
)

type InviteParticipantRequest struct {
	Id          string                               `json:"id" valid:"required,uuid"`
	AddressFrom string                               `json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Address     string                               `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	MemberType  member_deal_type_enum.MemberDealType `json:"memberType" valid:"required,range(0|2)"`
	MsgHash     string                               `json:"msgHash" valid:"required"`
	Sig         SignDto                              `json:"sig" valid:"required"`
	Exp         int64                                `json:"exp" valid:"required~ErrorTimestampNotPassed"`
}
