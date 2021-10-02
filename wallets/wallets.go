package wallets

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/mr-tron/base58"
	"github.com/ungame/go-bitcoin/keys"
	"golang.org/x/crypto/ripemd160"
	"log"
	"strings"
)

const (
	checksumLength = 4

	// https://en.bitcoin.it/wiki/List_of_address_prefixes
	version    = byte(0x00)
	wifVersion = byte(0x80)
)

type Wallet struct {
	k *keys.Keys
}

func New(k *keys.Keys) *Wallet {
	return &Wallet{k: k}
}

func (w *Wallet) GetPrivateKey() []byte {
	return w.k.GetPrivateKey()
}

func (w *Wallet) GetRipemd160() []byte {
	hash := sha256.Sum256(w.k.GetPublicKey())
	hasher := ripemd160.New()
	_, err := hasher.Write(hash[:])
	if err != nil {
		log.Panicln(err)
	}
	return hasher.Sum(nil)
}

// GetAddress implements: https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses
func (w *Wallet) GetAddress() []byte {
	hash160 := w.GetRipemd160()
	vHash160 := append([]byte{version}, hash160...)
	checksum := sum(vHash160)
	fullHash := append(vHash160, checksum...)
	encoded := base58.Encode(fullHash)
	return []byte(encoded)
}

func sum(versionedHash160 []byte) []byte {
	hash1 := sha256.Sum256(versionedHash160)
	hash2 := sha256.Sum256(hash1[:])
	return hash2[:checksumLength]
}

// WIF implements: https://developer.bitcoin.org/devguide/wallets.html#wallet-import-format-wif
func (w *Wallet) WIF() []byte {
	vPvtKey := append([]byte{wifVersion}, w.GetPrivateKey()...)
	hash1 := sha256.Sum256(vPvtKey)
	hash2 := sha256.Sum256(hash1[:])
	checksum := hash2[:checksumLength]
	wif := append(vPvtKey, checksum...)
	encoded := base58.Encode(wif)
	return []byte(encoded)
}

func (w *Wallet) String() string {
	return fmt.Sprintf(`
BTC WALLET:
	PrivateKey %s
	PublicKey  %s
	Ripemd160  %s
	Address    %s
	WIF        %s
`,
		strings.ToUpper(hex.EncodeToString(w.GetPrivateKey())),
		strings.ToUpper(hex.EncodeToString(w.k.GetPublicKey())),
		strings.ToUpper(hex.EncodeToString(w.GetRipemd160())),
		w.GetAddress(),
		w.WIF())
}
