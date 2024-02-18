package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCompleteOrganisationTransferTx(t *testing.T) {
	org1, toProfile, transfer := createRandomOrganisationTransfer(t)

	arg := CompleteOrganisationTransferTxParams{
		TransferId: transfer.ID,
		ToProfile:  toProfile.ID,
	}

	org2, err := testQueries.CompleteOrganisationTransferTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, org2)

	require.Equal(t, org1.ID, org2.ID)
	require.Equal(t, org1.Name, org2.Name)
	require.Equal(t, org1.Description, org2.Description)
	require.Equal(t, arg.ToProfile, org2.Owner)
	require.WithinDuration(t, org1.Createdat, org2.Createdat, time.Second)
}
