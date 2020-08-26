package models

//Operations An "Operation object" is a JSON object with information about an operation. Fields are:
type Operations struct {
	Valid      bool         `json:"valid"`
	Errors     string       `json:"errors"`
	Block      int          `json:"block"`
	Time       int          `json:"time"`
	OpBlock    int          `json:"opblock"`
	Maturation int          `json:"maturation"`
	OpType     int          `json:"optype"` // Options from 0 to 9
	Account    int          `json:"account"`
	Optxt      string       `json:"optxt"`
	Amount     float64      `json:"amount"`
	Fee        float64      `json:"fee"`
	Balance    float64      `json:"balance"`
	Senders    []OpSender   `json:"senders"`
	Receivers  []OpReciever `json:"recievers"`
	Changers   []OpChangers `json:"changers"`
}

//OpSender contains details of sender
type OpSender struct {
	Account    int        `json:"account"`
	NOperation int        `json:"n_operation"`
	Amount     float64    `json:"amount"`
	Payload    HexaString `json:"payload"`
}

//OpReciever contains details of the operation receiver
type OpReciever struct {
	Account int        `json:"account"`
	Amount  float64    `json:"amount"`
	Payload HexaString `json:"payload"`
}

//OpChangers contains details of changes made
type OpChangers struct {
	Account          int          `json:"account"`
	NOperation       int          `json:"n_operation"`
	NewEncPubkey     HexaString   `json:"new_enc_pubkey"`
	NewName          PascalCoin64 `json:"new_name"`
	NewType          int          `json:"new_type"`
	SellerAccount    int          `json:"seller_account"`
	AccountPrice     PascCurrency `json:"account_price"`
	LockedUntilBlock int          `json:"locked_until_block"`
	Fee              PascCurrency `json:"fee"`
}
