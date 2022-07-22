package helpers

import (
	"fmt"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateNewToken(length int) string {
	token := make([]rune, length)
	for i := range token {
		token[i] = letters[rand.Intn(len(letters))]
	}
	return string(token)
}

func PrintAndReturnPossibleError(err error, errorMessage string) error {
	if err != nil {
		fmt.Printf(errorMessage+" error %s", err)
		return err
	}
	return nil
}
