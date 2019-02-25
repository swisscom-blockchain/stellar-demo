package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	address := "GBD5HCK3EDWYPKSWXKQNIWB637INON5Y6J45HETVZXKFSWVMMJDSHO4Z"
	resp, err := http.Get("https://friendbot.stellar.org/?addr=" + address)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
