package main

import (
	"fmt"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
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

	tx, err := build.Transaction(
		build.SourceAccount{issuer.Address()},
		build.TestNetwork,
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.Payment(
			build.Destination{recipient.Address()},
			build.CreditAmount{"SCHF", issuer.Address(), "12"},
		),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(issuerSeed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successful Transaction:")
	fmt.Println("Ledger:", resp.Ledger)
	fmt.Println("Hash:", resp.Hash)
}
