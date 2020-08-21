package models

// PublicKey A "Public Key object" is a JSON object with information about a public key.
type PublicKey struct {
	Name      string `json:"name"`
	EncPubkey string `json:"enc_pubkey"`
	CanUse    bool   `json:"can_use"`
	B58pubkey string `json:"b58_pubkey"`
	EcNid     int    `json:"ec_nid"`
	X         string `json:"x"`
	Y         string `json:"y"`
}
