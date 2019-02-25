package main

import (
	"fmt"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"log"
)

func main() {

	issuerSeed := "SAQC3JLSSFYIUMIDUMIFQD4IPNQFJXJJUMDSTGNS5ALY3NVZ2XLMKDB3"
	recipientSeed := "SDRGLGG7KBNH6NFLGOU7FWEXZ4UAIA3WJDGUVLKL2QKYMAWRKAZOIW64"

	issuer, err := keypair.Parse(issuerSeed)
	if err != nil {
		log.Fatal(err)
	}
	recipient, err := keypair.Parse(recipientSeed)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := horizon.DefaultTestNetClient.LoadAccount(recipient.Address()); err != nil {
		panic(err)
	}

	passphrase := network.TestNetworkPassphrase

	tx, err := build.Transaction(
		build.Network{passphrase},
		build.SourceAccount{issuer.Address()},
		build.AutoSequence{horizon.DefaultTestNetClient},
		build.MemoText{"Payment Ref. 0001"},
		build.Payment(
			build.Destination{recipient.Address()},
			build.NativeAmount{"0.1"},
		),
	)

	if err != nil {
		panic(err)
	}

	txe, err := tx.Sign(issuerSeed)
	if err != nil {
		panic(err)
	}

	txeB64, err := txe.Base64()
	if err != nil {
		panic(err)
	}

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successful Transaction:")
	fmt.Println("Ledger:", resp.Ledger)
	fmt.Println("Hash:", resp.Hash)
}
