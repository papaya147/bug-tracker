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

	require.NotZero(t, org.Createdat)
	require.NotZero(t, org.Updatedat)

	return org
}

func TestCreateOrganisation(t *testing.T) {
	createRandomOrganisation(t)
}

func TestGetOrganisation(t *testing.T) {
	org1 := createRandomOrganisation(t)

	arg := org1.Owner
	org2, err := testQueries.GetOrganisation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, org2)

	require.Equal(t, org1.ID, org2.ID)
	require.Equal(t, org1.Name, org2.Name)
	require.Equal(t, org1.Description, org2.Description)
	require.Equal(t, org1.Owner, org2.Owner)
}

func TestUpdateOrganisation(t *testing.T) {
	org1 := createRandomOrganisation(t)

	arg := UpdateOrganisationParams{
		Name:        util.RandomString(10),
		Description: util.RandomString(100),
		Owner:       org1.Owner,
	}

	org2, err := testQueries.UpdateOrganisation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, org2)

	require.Equal(t, org1.ID, org2.ID)
	require.Equal(t, arg.Name, org2.Name)
	require.Equal(t, arg.Description, org2.Description)
	require.Equal(t, arg.Owner, org2.Owner)

	require.NotEqual(t, org1.Updatedat, org2.Updatedat)
}
