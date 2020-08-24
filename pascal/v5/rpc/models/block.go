package models

//Block is a JSON object with information about a Blockchain's block.
type Block struct {
	Block       int          `json:"block"`
	EncPubkey   HexaString   `json:"enc_pubkey"`
	Reward      PascCurrency `json:"reward"`
	RewardS     HexaString   `json:"reward_s"`
	Fee         PascCurrency `json:"fee"`
	Ver         int          `json:"ver"`
	VerA        int          `json:"ver_a"`
	Timestamp   int          `json:"timestamp"`
	Target      int          `json:"target"`
	Nonce       int          `json:"nonce"`
	Payload     string       `json:"payload"`
	Sbh         HexaString   `json:"sbh"`
	Oph         HexaString   `json:"oph"`
	Pow         HexaString   `json:"pow"`
	Operations  int          `json:"operations"`
	HashrateKhs int          `json:"hashratekhs"`
	Maturation  int          `json:"maturation"`
}
