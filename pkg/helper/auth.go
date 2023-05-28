package helper

import (
	"math/rand"
	"time"
)

const charset = "1234567890"

type Payload struct {
	Id        string      `json:"id"`
	ExpiredAt time.Time   `json:"expiredAt"`
	IssuedAt  time.Time   `json:"issuedAt"`
	Data      interface{} `json:"data"`
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// Generate random string from charset with defined length
func GenerateRandomString(length int) string {
	otp := make([]byte, length)
	for i := range otp {
		otp[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(otp)
}
