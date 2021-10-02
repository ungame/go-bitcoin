# Golang Bitcoin Protocols

## Usage

```bash
go get github/ungame/go-bitcoin@v1.0.0
```

## Bitcoin Wallets

```go
package main

import (
	"flag"
	"fmt"
	"github.com/ungame/go-bitcoin/keys"
	"github.com/ungame/go-bitcoin/wallets"
)

func main() {
	keyPair := keys.New()

	fmt.Println(keyPair)

	wallet := wallets.New(keyPair)

	fmt.Println(wallet)
}

```

#### Get Keys from Private Key in Hexdecimal

```go
package main

import (
	"flag"
	"fmt"
	"github.com/ungame/go-bitcoin/keys"
	"github.com/ungame/go-bitcoin/wallets"
)

func main() {
    	privateKey := "18e14a7b6a307f426a94f8114701e7c8e774e7f9a47e2c2035db29a206321725"

	keyPair := keys.FromPrivateKey([]byte(privateKey))

	fmt.Println(keyPair)

	wallet := wallets.New(keyPair)

	fmt.Println(wallet)
}

```

output

```bash
KEY PAIR:
        PrivateKey 18e14a7b6a307f426a94f8114701e7c8e774e7f9a47e2c2035db29a206321725
        PublicKey  0450863ad64a87ae8a2fe83c1af1a8403cb53f53e486d8511dad8a04887e5b23522cd470243453a299fa9e77237716103abc11a1df38855ed6f2ee187e9c582ba6


BTC WALLET:
        PrivateKey 18E14A7B6A307F426A94F8114701E7C8E774E7F9A47E2C2035DB29A206321725
        PublicKey  0450863AD64A87AE8A2FE83C1AF1A8403CB53F53E486D8511DAD8A04887E5B23522CD470243453A299FA9E77237716103ABC11A1DF38855ED6F2EE187E9C582BA6
        Ripemd160  010966776006953D5567439E5E39F86A0D273BEE
        Address    16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM
        WIF        5J1F7GHadZG3sCCKHCwg8Jvys9xUbFsjLnGec4H125Ny1V9nR6V
```


### References:

- https://developer.bitcoin.org/devguide/wallets.html#private-key-formats
- https://developer.bitcoin.org/devguide/wallets.html#public-key-formats
- https://developer.bitcoin.org/devguide/wallets.html#wallet-import-format-wif
- https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses
- https://en.bitcoin.it/wiki/List_of_address_prefixes
