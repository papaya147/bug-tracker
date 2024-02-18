package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateOrganisationTransferTx(t *testing.T) {
	org := createRandomOrganisation(t)
	profile2 := createRandomProfile(t)

	arg := CreateOrganisationTransferTxParams{
		FromProfile: org.Owner,
		ToEmail:     profile2.Email,
	}

	transfer, err := testQueries.CreateOrganisationTransferTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.Organisation, org.ID)
	require.Equal(t, transfer.Fromprofile, org.Owner)
	require.Equal(t, transfer.Toprofile, profile2.ID)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.Createdat)
}
