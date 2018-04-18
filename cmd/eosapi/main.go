package main

import (
	"bytes"
	"fmt"
	"log"
	"net/url"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
	"github.com/eoscanada/eos-go/system"
)

func main() {
	api := eos.New(&url.URL{Scheme: "http", Host: "cbillett.eoscanada.com"}, bytes.Repeat([]byte{0}, 32))
	//api := eos.New(&url.URL{Scheme: "http", Host: "localhost:8889"}, bytes.Repeat([]byte{0}, 32))

	// api.Debug = true

	keyBag := eos.NewKeyBag()
	for _, key := range []string{
		"5KQwrPbwdL6PhXujxW37FSSQZ1JiwsST4cqQzDeyXtP79zkvFD3",
		"5K7Ffo8LXHhbsxV48w3sZzo8UnaKX3z5iD5mvac1AfDhHXKs3ao",
		"5KE5hGNCAs1YvV74Ho14y1rV1DrnqZpTwLugS8QvYbKbrGAvVA1", // EOS71W8hvF43Eq6GQBRhuc5mvWKtknxzmb9NzNwPGpcEm2xAZaG8c
	} {
		if err := keyBag.Add(key); err != nil {
			log.Fatalln("Couldn't load private key:", err)
		}
	}

	api.SetSigner(keyBag)

	// Corresponding to the wallet, so we can sign on the live node.

	// resp, err := api.SetCode(AC("eosio"), "/home/abourget/build/eos/build/contracts/eosio.system/eosio.system.wasm", "/home/abourget/build/eos/build/contracts/eosio.system/eosio.system.abi")
	// if err != nil {
	// 	fmt.Println("ERROR calling SetCode:", err)
	// } else {
	// 	fmt.Println("RESP:", resp)
	// }

	resp, err := api.SignPushActions(
		system.NewNewAccount(AC("eosio"), AC("abourget"), ecc.MustNewPublicKey("EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV")),
	)
	if err != nil {
		fmt.Println("ERROR calling NewAccount:", err)
	} else {
		fmt.Println("RESP:", resp)
	}

	// walletAPI := eos.New(&url.URL{Scheme: "http", Host: "localhost:6667"}, bytes.Repeat([]byte{0}, 32))
	// walletAPI.Debug = true
	// api.SetSigner(eos.NewWalletSigner(walletAPI, "default"))

	// resp, err = api.SignPushActions(
	// 	system.NewNewAccount(AC("eosio"), AC("abourget"), ecc.MustNewPublicKey("EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV")),
	// )
	// if err != nil {
	// 	fmt.Println("ERROR calling NewAccount:", err)
	// } else {
	// 	fmt.Println("RESP:", resp)
	// }

}
