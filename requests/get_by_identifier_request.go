package requests

type GetByIdentifierRequest struct {
	Identifier string `json:"identifier" valid:"required~60304,validHex64~60303"`
}
