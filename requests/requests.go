package requests

type CreateAccount struct {
	Address       string   `json:"address"  validate:"required,validHex40"`
	State         int      `json:"state" validate:"required"`
	CurrencyCode  int      `json:"currencyCode" validate:"required,gte=0,lte=999"`
	JuridicalType int      `json:"juridicalType" validate:"required"`
	IdentityType  int      `json:"identityType" validate:"required"`
	Type          int      `json:"type" validate:"required"`
	Identifiers   []string `json:"identifiers"`
	Timestamp     int64    `json:"timestamp" validate:"required"`
	PublicKey     string   `json:"publicKey" validate:"required"`
	MsgHash       string   `json:"msgHash" validate:"required"`
	Sig           SignDto  `json:"sig" validate:"required,dive"`
}

type CreateBank struct {
	Address string `json:"address"  validate:"required,validHex40"`
	State   int    `json:"state" validate:"required"`
	Name    string `json:"name"`
	BIK     string `json:"bik"`
	MspId   string `json:"mspId"`
	IsOwner bool   `json:"isOwner"`
	Conf    string `json:"conf"`
}

type GetBank struct {
	MSPId   string `json:"mspId" validate:"required"`
	Address string `json:"address" validate:"required,validHex40"`
}

type SignDto struct {
	R string `json:"r" validate:"required,validHex64"`
	S string `json:"s" validate:"required,validHex64"`
	V int    `json:"v" validate:"required,gte=27,lte=28"`
}
