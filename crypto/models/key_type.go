package models

// KeyType holds different keytypes supported by PASC
type KeyType struct {
	Value uint16
}

// IsValid checks if KeyType is valid
func (k *KeyType) IsValid() bool {
	return *k == SECP256K1 ||
		*k == SECP384R1 ||
		*k == SECT283K1 ||
		*k == SECP521R1
}

// SECP256K1 keys based on SECP256K1 curve
var (
	SECP256K1 KeyType = KeyType{Value: 714}
	SECP384R1         = KeyType{Value: 715}
	SECT283K1         = KeyType{Value: 729}
	SECP521R1         = KeyType{Value: 716}
)
