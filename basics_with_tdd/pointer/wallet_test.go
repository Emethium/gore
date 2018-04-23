package pointer

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw succesfully", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw without necessary funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: Bitcoin(startingBalance)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, InsufficientFundsError)
	})
}
