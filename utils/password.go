package utils

import (
	"crypto/rand"
	"math/big"
)

func GenaratePassword() (string, int) {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^&*()_+"
	password := ""
	count := 1

	for {
		passwordInitLen, _ := rand.Int(rand.Reader, big.NewInt(12))
		passwordLen := passwordInitLen.Int64() + int64(9)
		password = GetRandomChar(passwordLen, charset)
		if ValidatePassword(password) {
			break
		}
		count++
	}
	return password, count
}
