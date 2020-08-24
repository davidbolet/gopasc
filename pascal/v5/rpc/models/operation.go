package models

//Operations An "Operation object" is a JSON object with information about an operation. Fields are:
type Operations struct {
	Valid      bool         `json:""`
	Errors     string       `json:""`
	Block      int          `json:""`
	Time       int          `json:""`
	OpBlock    int          `json:""`
	Maturation int          `json:""`
	OpType     int          `json:""` // Options from 0 to 9
	Account    int          `json:""`
	Optxt      string       `json:""`
	Amount     PascCurrency `json:""`
	Fee        PascCurrency `json:""`
	Balance    PascCurrency `json:""`
}
