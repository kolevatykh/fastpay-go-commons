package member_deal_type_enum

type MemberDealType int

const (
	Undefined MemberDealType = iota
	Participant
	Arbitrator
)

func (memberDealType MemberDealType) Str() string {
	return [...]string{"Undefined", "Participant", "Arbitrator"}[memberDealType]
}

func Parse(memberDealType string) MemberDealType {
	switch memberDealType {
	case "Participant":
		return Participant
	case "Deposit":
		return Arbitrator
	default:
		return Arbitrator
	}
}

func GetAllMemberDealTypes() []MemberDealType {
	return []MemberDealType{Undefined, Participant, Arbitrator}
}
