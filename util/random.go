package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "qwertyuiopasdfghjklzxcvbnm"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	value := max - min + 1
	return min + rand.Int63n(value)
}

// RandomString generates a random string of lenght n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}

	n := len(currencies)

	return currencies[rand.Intn(n)]
}
