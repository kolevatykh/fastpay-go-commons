package models

type AccountModel struct {
	Address                    string   `json:"address"`
	State                      int      `json:"state"`
	CurrencyCode               int      `json:"currencyCode"`
	JuridicalType              int      `json:"juridicalType"`
	BikBankSetterJuridicalType string   `json:"bikBankSetterJuridicalType"`
	IdentityType               int      `json:"identityType"`
	Owner                      string   `json:"owner"`
	Type                       int      `json:"type"`
	Identifiers                []string `json:"identifiers"`
	Encrypted                  bool     `json:"encrypted"`
	Created                    int64    `json:"created"`
	PublicKey                  string   `json:"publicKey"`
	DocType                    string   `json:"docType"`
}

type Bank struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	BIK         string `json:"bik"`
	State       int    `json:"state"`
	CreatedBy   string `json:"createdBy"`
	IsOwner     bool   `json:"isOwner"`
	Encrypted   bool   `json:"encrypted"`
	IsRegulator bool   `json:"isRegulator"`
	MSPId       string `json:"mspId"`
	Conf        string `json:"conf"`
	DocType     string `json:"docType"`
}
