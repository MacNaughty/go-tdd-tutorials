package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t1 testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t1.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t1 testing.TB, got error, want error) {
		t1.Helper()
		if got == nil {
			t1.Fatal("wanted an error but didn't get one")
		}

		if got != want {
			t1.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Deposit", func(t1 *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(40)

		assertBalance(t1, wallet, Bitcoin(40))
	})

	t.Run("Withdraw", func(t1 *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}

		wallet.Withdraw(Bitcoin(10))
		assertBalance(t1, wallet, Bitcoin(20))
	})

	t.Run("Withdraw insufficient funds", func(t1 *testing.T) {
		startingBalance := Bitcoin(55)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t1, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})

}
