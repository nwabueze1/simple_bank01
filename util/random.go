package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(size int) string {
	var stringBuilder strings.Builder
	alphabetSize := len(alphabets)

	for i := 0; i < size; i++ {
		character := alphabets[rand.Intn(alphabetSize)]
		stringBuilder.WriteByte(character)
	}

	return stringBuilder.String()
}

func GenerateRandomCurrency() string {
	currencies := []string{
		"USD",
		"CAD",
		"NGN",
		"GDP",
	}

	size := len(currencies)
	return currencies[rand.Intn(size)]
}

func GenerateRandomOwner() string {
	return RandomString(10)
}

func GenerateRandomAmount() int64 {
	return RandomInt(100, 1000)
}
