package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Test wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})

	t.Run("Test string", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		got := fmt.Sprint(wallet)
		want := "20 BTC"
		if got != want {
			t.Errorf("wanted %s, got %s", want, got)
		}
	})

	t.Run("Test withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		err := wallet.Withdraw(Bitcoin(30))
		if err == nil {
			t.Errorf("withdrawing %d from a %v. Err should not be nil", 30, wallet)
		}
		if err != ErrInsufficientFunds {
			t.Errorf("Error type is wrong, wanted %q, got %q", ErrInsufficientFunds, err)
		}
	})

}
