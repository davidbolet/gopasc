package models

// PascCurrency  Pascal Coin currency is a maximum 4 decimal number (ex. 12.1234). Decimal separator is a "." (dot).
type PascCurrency string

// HexaString String that contains an hexadecimal value (ex. "4423A39C"). An hexadecimal string is always an even character length..
type HexaString string

// PascalCoin64 represents an string limited to supported pascalcoin charset.
type PascalCoin64 string

// Account is a JSON object with information about a pascalcoint account.
type Account struct {
	Account int `json:"account"`
	EncPubkey HexaString `json:"enc_pubkey"`
	BalanceFloat float64 `json:"balance"`
	Balance PascCurrency `json:"balance_s"`
	NOperation int `json:"n_operation"`
	UpdatedB int `json:"updated_b"`
	UpdatedBActive int `json:"updated_b_active_mode"`
	UpdatedBPassive int `json:"updated_b_passive_mode"`
	State string `json:"state"`
	LockedUntilBlock int `json:"locked_until_block"`
	Price float64 `json:"price"`
	SellerAccount int `json:"seller_account"`
	PrivateSale bool `json:"private_sale"`
	NewEncPubkey HexaString `json:"new_enc_pubkey"`
	Name PascalCoin64 `json:"name"`
	Type uint16 `json:"type"`
	Seal string `json:"seal"`
}