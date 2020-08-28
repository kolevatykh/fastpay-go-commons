package invite_status_enum

type InviteStatus int

const (
	Undefined InviteStatus = iota
	Invited
	Rejected
	Accepted
)

func (inviteStatus InviteStatus) String() string {
	return [...]string{"Undefined", "Invited", "Rejected", "Accepted"}[inviteStatus]
}
