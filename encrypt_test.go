package weave

import (
	"testing"
)

func TestNewAccount(t *testing.T) {
	email := "john@example.com"
	name := "kismw365lo7emoxr3ohojgpild6lph4b"

	acct := NewAccount(email)

	if name != acct.String() {
		t.Errorf("Expected %s, got %s.\n", name, acct.String())
	}
}
