package main

import (
	"log"

	"github.com/stellar/go/keypair"
)

func main() {
	pair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(pair.Seed())
	log.Println(pair.Address())

}
