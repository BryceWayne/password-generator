package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"unicode"
)

func generatePassword(length int) (string, error) {
	const (
		lower    = "abcdefghijklmnopqrstuvwxyz"
		upper    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		digits   = "0123456789"
		symbols  = "!@#$%^&*()-_=+[]{}|;:,.<>?"
		allChars = lower + upper + digits + symbols
	)

	var password []rune
	var prevType string
	var currType string
	var count int

	for len(password) < length {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
		if err != nil {
			return "", err
		}
		char := rune(allChars[index.Int64()])
		switch {
		case unicode.IsLower(char):
			currType = "lower"
		case unicode.IsUpper(char):
			currType = "upper"
		case unicode.IsDigit(char):
			currType = "digit"
		default:
			currType = "symbol"
		}

		if currType == prevType {
			count++
			if count >= 3 {
				continue
			}
		} else {
			count = 1
		}

		password = append(password, char)
		prevType = currType
	}
	return string(password), nil
}

func main() {
	length := flag.Int("length", 12, "password length")
	flag.Parse()

	password, err := generatePassword(*length)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Generated password:", password)
}
