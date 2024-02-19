package db

import (
	"context"
	"testing"

	"github.com/papaya147/buggy/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomTeamMember(t *testing.T) {
	_, team := createRandomTeam(t)
	profile := createRandomProfile(t)

	arg := CreateTeamMemberParams{
		Team:    team.ID,
		Profile: profile.ID,
		Admin:   util.RandomBool(),
	}

	member, err := testQueries.CreateTeamMember(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, member)

	require.Equal(t, arg.Team, member.Team)
	require.Equal(t, arg.Profile, member.Profile)
	require.Equal(t, arg.Admin, member.Admin)

	require.NotZero(t, team.Createdat)
	require.NotZero(t, team.Updatedat)
}

func TestCreateTeamMember(t *testing.T) {
	createRandomTeamMember(t)
}
