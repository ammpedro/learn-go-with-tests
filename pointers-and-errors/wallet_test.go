package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("should handle deposits", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("should handle withdrawals with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("should handle withdrawals from insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(999))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, expectedBalance Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != expectedBalance {
		t.Errorf("got %s, but want %s", got, expectedBalance)
	}
}

func assertError(t *testing.T, actualError error, expectedError error) {
	t.Helper()
	if actualError == nil {
		t.Fatal("expecting an error, but did not get one")
	}

	if actualError != expectedError {
		t.Errorf("got %q, but want %q", actualError, expectedError)
	}
}

func assertNoError(t *testing.T, expectedError error) {
	t.Helper()
	if expectedError != nil {
		t.Fatalf("recevied error: %q, but did not expect one", expectedError.Error())
	}
}
