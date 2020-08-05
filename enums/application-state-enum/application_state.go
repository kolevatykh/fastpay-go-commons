package application_state_enum

type ApplicationState int

const (
	Undefined                  ApplicationState = 0
	Created                    ApplicationState = 1
	WaitingResponseFromArbiter ApplicationState = 2
	AwaitingConfirmation       ApplicationState = 3
	Published                  ApplicationState = 4
)

func (state ApplicationState) String() string {
	return [...]string{"Undefined", "Created", "WaitingResponseFromArbiter", "AwaitingConfirmation", "Published"}[state]
}
