package weave

import (
	"testing"
)

func TestNewAccount(t *testing.T) {
	emails := []string{
			"john@example.com",
			"john@doe.com",
			"jane@doe.com"}
	names := []string{
			"kismw365lo7emoxr3ohojgpild6lph4b",
			"7wohs32cngzuqt466q3ge7indszva4of",
			"vuuf3eqgloxpxmzph27f5a6ve7gzlrms"}

	for i, val := range emails {
		acct := NewAccount(val)

		if names[i] != acct.String() {
			t.Errorf("Expected %s, got %s.\n", names[i], acct.String())
		}
	}
}
