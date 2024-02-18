package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetOrganisationTransfersTx(t *testing.T) {
	org, _, _ := createRandomOrganisationTransfer(t)

	transfers, err := testQueries.GetOrganisationTransfersTx(context.Background(), org.Owner)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
}
