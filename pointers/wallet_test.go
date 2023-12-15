package wallet

import "testing"

func Testwallet(t *testing.T) {
	wallet := wallet{}

	wallet.deposit(10)

	got := wallet.balance()
	want := 10
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
