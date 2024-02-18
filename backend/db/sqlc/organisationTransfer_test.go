package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomOrganisationTransfer(t *testing.T) {
	org := createRandomOrganisation(t)
	profile := createRandomProfile(t)

	arg := CreateOrganisationTransferParams{
		Organisation: org.ID,
		Fromprofile:  org.Owner,
		Toprofile:    profile.ID,
	}

	transfer, err := testQueries.CreateOrganisationTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.Organisation, org.ID)
	require.Equal(t, transfer.Fromprofile, org.Owner)
	require.Equal(t, transfer.Toprofile, profile.ID)

	require.NotZero(t, transfer.Createdat)
}

func TestCreateOrganisationTransfer(t *testing.T) {
	createRandomOrganisationTransfer(t)
}
