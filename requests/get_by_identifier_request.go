package requests

type GetByIdentifierRequest struct {
	Identifier string `json:"identifier" valid:"required~ErrorIdentifierNotPassed,validHex64~ErrorIdentifierNotFolowingRegex"`
}
