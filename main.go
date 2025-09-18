package main

import (
	"fmt"
)

type account struct {
	login    string
	password string
	url      string
}

func main() {
	userLogin := promptUserData("Enter your login")
	userPassword := promptUserData("Enter you password")
	userUrl := promptUserData("Enter you url")

	account := account{
		login:    userLogin,
		password: userPassword,
		url:      userUrl,
	}

	outputUserData(account)
}

func promptUserData(prompt string) string {
	var userAnswer string
	fmt.Print(prompt + ": ")
	fmt.Scan(&userAnswer)

	return userAnswer
}

func outputUserData(userAccount account) {
	message := fmt.Sprintf("Dear %v, your password is %v and it's reference to https://%v.com", userAccount.login, userAccount.password, userAccount.url)
	fmt.Println(message)
}

func reverse(array *[5]int) {
	for index, value := range *array {
		(*array)[len(array)-1-index] = value
	}
}
