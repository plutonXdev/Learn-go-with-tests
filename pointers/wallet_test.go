package wallet

import "testing"

func TestWallet(t *testing.T) {
	Wallet := Wallet{}

	Wallet.Deposit(10)

	got := Wallet.Balance()
	want := 10
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
