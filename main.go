package main

import (
	"flag"
	"fmt"
	"github.com/ungame/go-bitcoin/keys"
	"github.com/ungame/go-bitcoin/wallets"
)

var (
	privateKey string
)

func init() {
	flag.StringVar(&privateKey, "pvtKey", "18e14a7b6a307f426a94f8114701e7c8e774e7f9a47e2c2035db29a206321725", "private key in hexadecimal")
	flag.Parse()
}

func main() {
	keyPair := keys.FromPrivateKey([]byte(privateKey)) //keys.New()

	fmt.Println(keyPair)

	wallet := wallets.New(keyPair)

	fmt.Println(wallet)
}
