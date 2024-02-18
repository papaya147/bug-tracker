package db

import (
	"context"
	"testing"

	"github.com/papaya147/buggy/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomOrganisationTransfer(t *testing.T) (Organisation, Profile, Organisationtransfer) {
	org := createRandomOrganisation(t)
	profile := createRandomProfile(t)

	arg := CreateOrganisationTransferParams{
		ID:           util.RandomUuid(),
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

	return org, profile, transfer
}

func TestCreateOrganisationTransfer(t *testing.T) {
	createRandomOrganisationTransfer(t)
}

func TestGetActiveOrganisationTransfer(t *testing.T) {
	org, _, transfer1 := createRandomOrganisationTransfer(t)

	transfer2, err := testQueries.GetActiveOrganisationTransfer(context.Background(), org.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.Organisation, transfer2.Organisation)
	require.Equal(t, transfer1.Fromprofile, transfer2.Fromprofile)
	require.Equal(t, transfer1.Toprofile, transfer2.Toprofile)
}

func TestDeleteOrganisationTransfer(t *testing.T) {
	org, _, transfer1 := createRandomOrganisationTransfer(t)

	arg := DeleteOrganisationTransferParams{
		ID:          transfer1.ID,
		Fromprofile: org.Owner,
	}

	transfer2, err := testQueries.DeleteOrganisationTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.Organisation, transfer2.Organisation)
	require.Equal(t, transfer1.Fromprofile, transfer2.Fromprofile)
	require.Equal(t, transfer1.Toprofile, transfer2.Toprofile)
}
