package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"unicode"
)

const PASSWORD_LENGTH = 13

type userAccount struct {
	login    string
	password string
	url      string
}

func (account userAccount) outputUserData() string {
	return fmt.Sprintf(
		"Dear user %s, your password is %s and it's reference to %s\n", CapitalizeWord(account.login), account.password, account.url)
}

func (account *userAccount) generatePassword(n int) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, n)

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	account.password = string(result)
}

func newAccount(userLogin, userPassword, userUrl string) (*userAccount, error) {
	_, err := url.ParseRequestURI(userUrl)

	if err != nil {
		return nil, errors.New("invalid URL")
	}

	if userLogin == "" {
		return nil, errors.New("login cannot be empty")
	}

	acc := &userAccount{
		login:    userLogin,
		url:      userUrl,
		password: userPassword,
	}

	if userPassword == "" {
		acc.generatePassword(PASSWORD_LENGTH)
	}

	return acc, nil
}

func main() {
	userLogin := promptUserData("Enter your login")
	userPassword := promptUserData("Enter your password")
	userUrl := promptUserData("Enter you url (yandex/google)")

	account, err := newAccount(userLogin, userPassword, userUrl)

	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}

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
