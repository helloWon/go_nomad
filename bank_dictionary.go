package main

import (
	"fmt"
	"hello/mydict"
)

func bank_dictionary() {
	// account := accounts.NewAccount("nico")

	// account.Deposit(10)
	// fmt.Println(account.Balance())

	// // err := account.Withdraw(20)
	// // if err != nil {
	// // 	log.Fatalln(err)
	// // }

	// err := account.Withdraw(10)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(account.String())

	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	fmt.Println(dictionary["first"])

	word := "hello"
	def := "Greeting"
	err = dictionary.Add(word, def)
	if err != nil {
		fmt.Println(err)
	}
	hello, err2 := dictionary.Search("hello")
	fmt.Println(hello)
	if err2 != nil {
		fmt.Println(err2)
	}

	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	// Update
	err = dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ = dictionary.Search(baseWord)
	fmt.Println(word)

	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, _ = dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
