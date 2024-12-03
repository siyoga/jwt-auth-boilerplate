package random

import (
	"encoding/hex"
	"math/rand"
	"time"

	def "github.com/siyoga/jwt-auth-boilerplate/internal/adapter"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz1234567890"
	allChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var _ def.RandomAdapter = (*adapter)(nil)

type (
	adapter struct {
	}
)

func NewAdapter() def.RandomAdapter {
	return &adapter{}
}

func (a *adapter) RandomToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	token := hex.EncodeToString(bytes)
	return token, nil
}

func (a *adapter) RandomString(length int) string {
	return a.randomString(length)
}

func (a *adapter) RandomIntn(n int) int {
	return rand.Intn(n)
}

func (a *adapter) RandomStringWithTimeNanoSeed(length int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = allChars[r.Intn(len(allChars))]
	}
	return string(result)
}

func (a *adapter) randomString(length int) string {
	res := make([]byte, length)
	for i := 0; i < length; i++ {
		res[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(res)
}
