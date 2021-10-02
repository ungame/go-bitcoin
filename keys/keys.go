package keys

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/haltingstate/secp256k1-go"
	"log"
	"math/big"
)

type Keys struct {
	privateKey *ecdsa.PrivateKey
}

func New() *Keys {
	curve := elliptic.P256()

	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panicln(err)
	}

	return &Keys{privateKey: privateKey}
}

func FromPrivateKey(privateKey []byte) *Keys {
	var pvtKey ecdsa.PrivateKey
	var pvtHex big.Int
	pvtKey.D, _ = pvtHex.SetString(string(privateKey), 16)
	pvtKey.PublicKey.Curve = elliptic.P256()
	pvtKey.PublicKey.X, pvtKey.PublicKey.Y = pvtKey.PublicKey.Curve.ScalarBaseMult(pvtKey.D.Bytes())
	return &Keys{privateKey: &pvtKey}
}

// GetPrivateKey implements: https://developer.bitcoin.org/devguide/wallets.html#private-key-formats
func (k *Keys) GetPrivateKey() []byte {
	return k.privateKey.D.Bytes()
}

// GetPublicKey implements: https://developer.bitcoin.org/devguide/wallets.html#public-key-formats
func (k *Keys) GetPublicKey() []byte {
	return secp256k1.UncompressedPubkeyFromSeckey(k.privateKey.D.Bytes())
}

func (k *Keys) String() string {
	return fmt.Sprintf(`
KEY PAIR:
	PrivateKey %s
	PublicKey  %s
`, hex.EncodeToString(k.GetPrivateKey()), hex.EncodeToString(k.GetPublicKey()))
}
