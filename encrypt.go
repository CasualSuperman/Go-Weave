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
	email = strings.ToLower(email)

	buf := make([]byte, sha1.Size)

	sha := sha1.New()
	sha.Write([]byte(email))
	buf = sha.Sum(buf)

	result := make([]byte, base32.StdEncoding.EncodedLen(len(buf)))
	base32.StdEncoding.Encode(result, buf)

	name := strings.ToLower(string(result[32:]))

	return Account{name}
}

func (a Account) String() string {
	return a.username
}
