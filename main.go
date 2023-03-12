package main

import (
	"fmt"
	"hello/accounts"
	"log"
)

func main() {
	account := accounts.NewAccount("nico")

	account.Deposit(10)
	fmt.Println(account.Balance())

	// err := account.Withdraw(20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err := account.Withdraw(10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.String())
}
