package weave

import (
	"testing"
)

func TestNewAccount(t *testing.T) {
	data := []struct {
		Email, Hash string
	}{
		{"john@doe.com", "7wohs32cngzuqt466q3ge7indszva4of"},
		{"jane@doe.com", "vuuf3eqgloxpxmzph27f5a6ve7gzlrms"},
		{"john@example.com", "kismw365lo7emoxr3ohojgpild6lph4b"},
	}

	for _, user := range data {
		acct := NewAccount(user.Email)

		if user.Hash != acct.String() {
			t.Errorf("Expected %s, got %s.\n", user.Hash, acct.String())
		}
	}
}
