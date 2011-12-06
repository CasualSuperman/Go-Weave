package weave

import (
	"crypto/sha1"
	"encoding/base32"
	"strings"
)

type Account struct {
	username string
}

func NewAccount(email string) Account {
	name := strings.ToLower(email)

	username := make([]byte, sha1.Size)

	sha := sha1.New()
	sha.Write([]byte(name))
	username = sha.Sum(username)

	result := make([]byte, base32.StdEncoding.EncodedLen(len(username)))
	base32.StdEncoding.Encode(result, username)

	name = strings.ToLower(string(result[32:]))

	return Account{name}
}

func (a Account) String() string {
	return a.username
}
