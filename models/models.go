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
