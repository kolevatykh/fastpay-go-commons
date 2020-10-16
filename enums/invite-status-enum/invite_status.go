package invite_status_enum

type InviteStatus int

const (
	Undefined InviteStatus = iota
	Invited
	Rejected
	Accepted
)

func (inviteStatus InviteStatus) Str() string {
	return [...]string{"Undefined", "Invited", "Rejected", "Accepted"}[inviteStatus]
}
