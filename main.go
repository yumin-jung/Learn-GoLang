package main

import (
	"fmt"

	"github.com/yumin-jung/Learn-GoLang/mydict"
)

func main() {
	// Accounts

	// account := accounts.NewAccount("yumin")
	// account.Deposit(10)
	// errWithdraw := account.Withdraw(5)
	// if errWithdraw != nil {
	// 	fmt.Println(errWithdraw)
	// }
	// fmt.Println(account)

	// Dict
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "first")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
