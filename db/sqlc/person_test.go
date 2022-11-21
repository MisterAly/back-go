package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/MisterAly/weesir/utils"
	"github.com/stretchr/testify/require"
)

func createRandomPerson(t *testing.T) Person {
	arg := CreatePersonParams {
		Name: utils.RandomName(),
		Document: utils.RandomDocument(),
		Phone: sql.NullString{},
	}

	person, err := testQueries.CreatePerson(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, person)

	require.Equal(t, arg.Name, person.Name)
	require.Equal(t, arg.Document, person.Document)

	require.NotZero(t, person.ID)
	require.NotZero(t, person.Document)

	return person
}

func TestCreatePerson(t *testing.T) {
	createRandomPerson(t)
}

func TestGetPerson(t *testing.T) {
	person1 :=	createRandomPerson(t)
	person2, err := testQueries.GetPerson(context.Background(), person1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, person2)

	require.Equal(t, person1.ID, person2.ID)
	require.Equal(t, person1.Name, person2.Name)
	require.Equal(t, person1.Document, person2.Document)
	require.Equal(t, person1.Phone, person2.Phone)
	require.WithinDuration(t, person1.CreatedAt, person2.CreatedAt, time.Second)
}

func TestUpdatePerson(t *testing.T) {
	person1 := createRandomPerson(t)

	arg := UpdatePersonParams{
		Name: person1.Name,
		Document: utils.RandomDocument(),
		Phone: sql.NullString{},
	}

	person2, err := testQueries.UpdatePerson(context.Background(), arg)
	
	require.NoError(t, err)
	require.NotEmpty(t, person2)

	require.Equal(t, person1.ID, person2.ID)
	require.Equal(t, person1.Name, person2.Name)
	require.Equal(t, arg.Document, person2.Document)
	require.Equal(t, person1.Phone, person2.Phone)
	require.WithinDuration(t, person1.CreatedAt, person2.CreatedAt, time.Second)
}

func TestDeletePerson(t *testing.T) {
	person1 := createRandomPerson(t)
	err := testQueries.DeletePerson(context.Background(), person1.ID)
	require.NoError(t, err)

	person2, err := testQueries.GetPerson(context.Background(), person1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, person2)
}

func TestListPerson(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPerson(t)
	}

	arg := ListPersonParams {
		Limit: 5,
		Offset: 5,
	}

	people, err := testQueries.ListPerson(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, people, 5)

	for _, person := range people {
		require.NotEmpty(t, person)
	}
}