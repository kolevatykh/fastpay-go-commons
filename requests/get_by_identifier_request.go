package requests

type GetByIdentifierRequest struct {
	Identifier string `json:"identifier" validate:"required,validHex64"`
}
