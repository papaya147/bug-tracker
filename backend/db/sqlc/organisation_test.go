package db

import (
	"context"
	"testing"

	"github.com/papaya147/buggy/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomOrganisation(t *testing.T) Organisation {
	profile := createRandomProfile(t)

	arg := CreateOrganisationParams{
		ID:          util.RandomUuid(),
		Name:        util.RandomString(10),
		Description: util.RandomString(100),
		Owner:       profile.ID,
	}

	org, err := testQueries.CreateOrganisation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, org)

	require.Equal(t, arg.ID, org.ID)
	require.Equal(t, arg.Name, org.Name)
	require.Equal(t, arg.Description, org.Description)
	require.Equal(t, arg.Owner, org.Owner)

	return org
}

func TestCreateOrganisation(t *testing.T) {
	createRandomOrganisation(t)
}
