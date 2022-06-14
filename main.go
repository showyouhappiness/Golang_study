package main

import (
	"go_test/src/bankapi"
	"log"
	"net/http"
)

func main() {
	bankapi.Login()
	http.HandleFunc("/statement", bankapi.Statement)
	http.HandleFunc("/deposit", bankapi.Deposit)
	http.HandleFunc("/withdraw", bankapi.Withdraw)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
