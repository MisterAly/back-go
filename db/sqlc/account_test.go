package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/MisterAly/weesir/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T, person Person) Account{
	arg := CreateAccountParams {
		PersonID: person.ID,
		Amount: utils.RandomMoney(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.PersonID, account.PersonID)
	require.Equal(t, arg.Amount, account.Amount)

	require.NotZero(t, account.PersonID)
	require.NotZero(t, account.Amount)

	return account
}

func TestCreateAccount(t *testing.T) {
	person := createRandomPerson(t)
	createRandomAccount(t, person)
}

func TestGetAccount(t *testing.T) {
	person := createRandomPerson(t)
	account1 := createRandomAccount(t, person)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.PersonID, account2.PersonID)
	require.Equal(t, account1.Amount, account2.Amount)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	person := createRandomPerson(t)
	account1 := createRandomAccount(t, person)

	arg := UpdateAccountParams{
		ID: account1.ID,
		Amount: utils.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.PersonID, account2.PersonID)
	require.Equal(t, arg.Amount, account2.Amount)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	person := createRandomPerson(t)
	account1 := createRandomAccount(t, person)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccount(t *testing.T) {
	person := createRandomPerson(t)
	for i := 0; i < 10;  i++ {
		createRandomAccount(t, person)
	}

	arg := ListAccountParams{
		PersonID: person.ID,
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, arg.PersonID, account.PersonID)
	}
}