package models

//DecryptResult holds the value returned by payloaddecrypt
type DecryptResult struct {
	Result           bool       `json:"result"`
	EncPayload       HexaString `json:"enc_payload"`
	UnEncPayload     string     `json:"unenc_payload"`
	UnEncHexaPayload HexaString `json:"unenc_hexpayload"`
	PayloadMethod    string     `json:"payload_method"`
	EncPubkey        HexaString `json:"enc_pubkey"`
	Pwd              string     `json:"pwd"`
}
