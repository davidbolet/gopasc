package models

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/akamensky/base58"
)

// WalletKey represents the structure of a key pair (public/private)
type WalletKey struct {
	PrivateKey []byte
	Public     *PublicKey
}

// PublicKey represents the information required for a public key
type PublicKey struct {
	KeyType   KeyType
	ECDSA     ecdsa.PublicKey
	B58pubkey string
	EncPubkey string
}

// NewPublicKey generates a new public key from KeyType and given ecdsa public key
func NewPublicKey(keyType KeyType, ECDSA *ecdsa.PublicKey) *PublicKey {
	var pascPublicKey PublicKey = PublicKey{KeyType: keyType, ECDSA: *ECDSA}
	pascPublicKey.B58pubkey = base58.Encode()
}

func checksum(public ecdsa.PublicKey) uint32 {
	binary.Write()
	public.TypeId,
	public.Bytes(),
	public.Y.Bytes(),
	toChecksum := utils.Serialize(&serialized)

	hash := sha256.Sum256(toChecksum)
	return binary.LittleEndian.Uint32(hash[:4])
}

// FromPrivateKey Returns a WalletKey object from a Private key string
func FromPrivateKey(privateKey string, keyType KeyType) (*WalletKey, error) {
	pkBytes := []byte(privateKey)
	curve, err := getCurveFromKeyType(keyType)
	if err != nil {
		return nil, err
	}
	x, y := curve.ScalarBaseMult(pkBytes)
	var ecdsaPublicKey ecdsa.PublicKey = ecdsa.PublicKey{curve, x, y}
	return &WalletKey{PrivateKey: pkBytes, Public: NewPublicKey(keyType, &ecdsaPublicKey)}, nil
}

func getCurveFromKeyType(keyType KeyType) (elliptic.Curve, error) {
	switch keyType {
	case SECP256K1:
		return elliptic.P256(), nil
	case SECP384R1:
		return elliptic.P384(), nil
	case SECT283K1:
		return nil, errors.New("Curve not implemented")
	case SECP521R1:
		return elliptic.P521(), nil
	}
	return nil, fmt.Errorf("Unknown curve %d", keyType.Value)
}

// GenerateWalletKey generates a new WalletKey object with a random Keypair
func GenerateWalletKey(keyType KeyType) (*WalletKey, error) {
	curve, err := getCurveFromKeyType(keyType)
	if err != nil {
		return nil, err
	}
	priv, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	var publicKey ecdsa.PublicKey = ecdsa.PublicKey{curve, x, y}

	return &WalletKey{PrivateKey: priv, Public: NewPublicKey(keyType, &publicKey)}, nil
}

// GetEcdsaPrivateKey generates the corresponding ecdsa.PrivateKey object
func (wk *WalletKey) GetEcdsaPrivateKey() *ecdsa.PrivateKey {
	return &ecdsa.PrivateKey{
		D:         big.NewInt(0).SetBytes(wk.PrivateKey),
		PublicKey: wk.Public.ECDSA,
	}
}
