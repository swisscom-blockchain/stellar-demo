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

	stableCoinChf := build.CreditAsset("SCHF", issuer.Address())

	tx, err := build.Transaction(
		build.SourceAccount{recipient.Address()},
		build.AutoSequence{horizon.DefaultTestNetClient},
		build.TestNetwork,
		build.Trust(stableCoinChf.Code, stableCoinChf.Issuer, build.Limit("100")),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(recipientSeed)
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

/*

tx base64: AAAAADZY/nWY0gx6beMpf4S8Ur0qHsjA8fbFtBzBx1cbQzHwAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAYAAAABQ0hGAAAAAABH04lbIO2Hqla6oNRYPt/Q1ze48nnTknXN1FlarGJHIwAAAAA7msoAAAAAAAAAAAEbQzHwAAAAQPUInbFCwn+bdnuJDZYesNO0W9LI83JkSGflMahND5Q6I0Xl0F1JVSqbi0lVOwQE8xAtirNkn7aLRzWcDSEKYQw=

*/
