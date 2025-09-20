package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

const passwordLength = 13

type userAccount struct {
	login    string
	password string
	url      string
}

func (account userAccount) outputUserData() string {
	return fmt.Sprintf(
		"Dear user %s, your password is %s and it's reference to https://%s.com\n", CapitalizeWord(account.login), account.password, account.url)
}

func (account *userAccount) generatePassword(n int) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, n)

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	account.password = string(result)
}

func main() {
	userLogin := promptUserData("Enter your login")
	userUrl := promptUserData("Enter you url")

	account := userAccount{
		login: userLogin,
		url:   userUrl,
	}
	account.generatePassword(passwordLength)

	result := account.outputUserData()
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

func CapitalizeWord(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(strings.ToLower(s))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
