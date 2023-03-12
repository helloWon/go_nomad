package main

import (
	"fmt"
	"hello/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
