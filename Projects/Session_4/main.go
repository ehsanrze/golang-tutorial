package main

import (
	"fmt"
	"github.com/ehsanrze/golang-tutorial/Session_4/user"
)

func main() {

	var name string
	var age uint
	var email string
	var phone string

	_, err := fmt.Scanln(&name, &age, &email, &phone)
	if err != nil {
		panic(err)
	}

	newUser, err := user.New(name, age, email, phone, user.Address{})
	if err != nil {
		panic(err)
	}

	newUser.Log()
	err = newUser.Write()
	if err != nil {
		panic(err)
	}

}
