package weave

import (
	"crypto/sha1"
	"encoding/base32"
	"strings"
)

type Account struct {
	username string
}

// Creates a new account and computes the hash of the username
func NewAccount(email string) Account {
	// Lowercase
	email = strings.ToLower(email)

	// sha1
	sha := sha1.New()
	sha.Write([]byte(email))
	buf := sha.Sum(nil)

	// Base32 encode
	result := make([]byte, base32.StdEncoding.EncodedLen(len(buf)))
	base32.StdEncoding.Encode(result, buf)

	// Lowercase
	name := strings.ToLower(string(result))

	return Account{name}
}

func (a Account) Hash() string {
	return a.username
}
