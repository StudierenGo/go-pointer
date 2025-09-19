package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type userAccount struct {
	login    string
	password string
	url      string
}

func main() {
	userLogin := promptUserData("Enter your login")
	userPassword := promptUserData("Enter you password")
	userUrl := promptUserData("Enter you url")

	account := userAccount{
		login:    userLogin,
		password: userPassword,
		url:      userUrl,
	}

	result := outputUserData(account)
	fmt.Println(result)
}

func promptUserData(prompt string) string {
	fmt.Print(prompt + ": ")

	reader := bufio.NewReader(os.Stdin)
	userAnswer, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	userAnswer = strings.TrimSpace(userAnswer)
	return userAnswer
}

func outputUserData(account userAccount) string {
	return fmt.Sprintf(
		"Dear user %s, your password is %s and it's reference to https://%s.com\n", CapitalizeWord(account.login), account.password, account.url)
}

func CapitalizeWord(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(strings.ToLower(s))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
