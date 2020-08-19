package requests

type GetByIdentifierRequest struct {
	Identifier string `json:"identifier" valid:"required,validHex64"`
}
