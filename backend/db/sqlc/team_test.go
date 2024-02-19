package db

import (
	"context"
	"testing"

	"github.com/papaya147/buggy/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomTeam(t *testing.T) (Organisation, Team) {
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

	return org, team
}

func TestCreateTeam(t *testing.T) {
	createRandomTeam(t)
}

func TestGetOrganisationTeams(t *testing.T) {
	org, team := createRandomTeam(t)

	arg := org.ID

	teams, err := testQueries.GetOrganisationTeams(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, teams)

	require.Contains(t, teams, team)
}

func TestUpdateTeam(t *testing.T) {
	org, team1 := createRandomTeam(t)

	arg := UpdateTeamParams{
		Name:         util.RandomString(10),
		Description:  util.RandomString(100),
		ID:           team1.ID,
		Organisation: org.ID,
	}

	team2, err := testQueries.UpdateTeam(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, team2)

	require.Equal(t, arg.ID, team2.ID)
	require.Equal(t, arg.Name, team2.Name)
	require.Equal(t, arg.Description, team2.Description)
	require.Equal(t, arg.Organisation, team2.Organisation)

	require.NotZero(t, team2.Createdat)
	require.NotZero(t, team2.Updatedat)
}

func TestGetTeamOrganisation(t *testing.T) {
	org, team1 := createRandomTeam(t)

	arg := team1.ID

	orgId, err := testQueries.GetTeamOrganisation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orgId)

	require.Equal(t, org.ID, orgId)
}
