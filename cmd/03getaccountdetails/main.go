package main

import (
	"fmt"
	"log"

	"github.com/stellar/go/clients/horizon"
)

func main() {

	address := "GB2Q3KZPUAUPOKKTBJBB5GNACRLV66HCDKCHBK2DCJE7UWNCPVCGX2DX"

	account, err := horizon.DefaultTestNetClient.LoadAccount(address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balances for account:", address)

	for _, balance := range account.Balances {
		log.Println(balance)
	}
}
