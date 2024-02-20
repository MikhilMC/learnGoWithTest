package pointers

import (
	// "fmt"
	"testing"
)

func TestWalllet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()

	// fmt.Printf("address of balance in the test is %p \n", &wallet.balance)

	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
