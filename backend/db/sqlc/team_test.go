package db

import (
	"context"
	"testing"

	"github.com/papaya147/buggy/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomTeam(t *testing.T) {
	org := createRandomOrganisation(t)

	arg := CreateTeamParams{
		ID:           util.RandomUuid(),
		Name:         util.RandomString(10),
		Description:  util.RandomString(100),
		Organisation: org.ID,
	}

	team, err := testQueries.CreateTeam(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, team)

	require.Equal(t, arg.ID, team.ID)
	require.Equal(t, arg.Name, team.Name)
	require.Equal(t, arg.Description, team.Description)
	require.Equal(t, arg.Organisation, team.Organisation)

	require.NotZero(t, team.Createdat)
	require.NotZero(t, team.Updatedat)
}

func TestCreateTeam(t *testing.T) {
	createRandomTeam(t)
}
