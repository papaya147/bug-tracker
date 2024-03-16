package db

import (
	"context"
	"testing"
	"time"

	"github.com/papaya147/buggy/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomTeamMember(t *testing.T) (Organisation, Team, Teammember) {
	org, team := createRandomTeam(t)
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

	return org, team, member
}

func TestCreateTeamMember(t *testing.T) {
	createRandomTeamMember(t)
}

func TestGetTeamMember(t *testing.T) {
	_, team, member1 := createRandomTeamMember(t)

	arg := GetTeamMemberParams{
		Team:    team.ID,
		Profile: member1.Profile,
	}

	member2, err := testQueries.GetTeamMember(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, member2)

	require.Equal(t, member1.Team, member2.Team)
	require.Equal(t, member1.Profile, member2.Profile)
	require.Equal(t, member1.Admin, member2.Admin)
	require.WithinDuration(t, member1.Createdat, member2.Createdat, time.Second)
	require.WithinDuration(t, member1.Updatedat, member2.Updatedat, time.Second)
}

func TestGetAllTeamMembers(t *testing.T) {
	_, team, _ := createRandomTeamMember(t)

	arg := team.ID

	members, err := testQueries.GetAllTeamMembers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, members)
}

func TestUpdateTeamMember(t *testing.T) {
	_, team, member1 := createRandomTeamMember(t)

	arg := UpdateTeamMemberParams{
		Admin:   util.RandomBool(),
		Team:    team.ID,
		Profile: member1.Profile,
	}

	member2, err := testQueries.UpdateTeamMember(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, member2)

	require.Equal(t, arg.Team, member2.Team)
	require.Equal(t, arg.Profile, member2.Profile)
	require.Equal(t, arg.Admin, member2.Admin)
}

func TestGetTeams(t *testing.T) {
	_, _, member := createRandomTeamMember(t)

	arg := member.Profile

	teams, err := testQueries.GetTeams(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, teams)
}
