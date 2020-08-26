package models

//RawOperations A "Raw operations object" is a JSON object with information about a signed operation made by "signsendto" or "signchangekey"
type RawOperations struct {
	Operations    int        `json:"operations"`
	Amount        float64    `json:"amount"`
	Fee           float64    `json:"fee"`
	Rawoperations HexaString `json:"rawoperations"`
}
