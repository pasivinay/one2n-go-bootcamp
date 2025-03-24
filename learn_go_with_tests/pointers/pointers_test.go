package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10.0))
	got := wallet.Balance()
	want := Bitcoin(10.0)

	if got != want {
		t.Errorf("Expected Balance: %v, Fetched Balance:%v", want, got)
	}
}
