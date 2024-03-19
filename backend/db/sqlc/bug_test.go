package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/papaya147/buggy/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomBug(t *testing.T) (Bug, Team, Teammember) {
	_, team, member := createRandomTeamMember(t)

	arg := CreateBugParams{
		ID:                util.RandomUuid(),
		Name:              util.RandomString(10),
		Description:       util.RandomString(100),
		Priority:          BugpriorityLOW,
		Assignedto:        team.ID,
		Assignedbyprofile: member.Profile,
		Assignedbyteam:    member.Team,
	}

	bug, err := testQueries.CreateBug(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bug)
	require.Equal(t, arg.ID, bug.ID)
	require.Equal(t, arg.Name, bug.Name)
	require.Equal(t, arg.Description, bug.Description)
	require.Equal(t, BugstatusPENDING, bug.Status)
	require.Equal(t, arg.Priority, bug.Priority)
	require.Equal(t, arg.Assignedto, bug.Assignedto)
	require.Equal(t, arg.Assignedbyprofile, bug.Assignedbyprofile)
	require.Equal(t, arg.Assignedbyteam, bug.Assignedbyteam)
	require.NotZero(t, bug.Completed)
	require.NotZero(t, bug.Createdat)
	require.NotZero(t, bug.Updatedat)

	return bug, team, member
}

func TestCreateBug(t *testing.T) {
	createRandomBug(t)
}

func TestGetBug(t *testing.T) {
	bug1, _, _ := createRandomBug(t)

	bug2, err := testQueries.GetBug(context.Background(), bug1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, bug2)
	require.Equal(t, bug1.ID, bug2.ID)
	require.Equal(t, bug1.Name, bug2.Name)
	require.Equal(t, bug1.Description, bug2.Description)
	require.Equal(t, bug1.Status, bug2.Status)
	require.Equal(t, bug1.Priority, bug2.Priority)
	require.Equal(t, bug1.Assignedto, bug2.Assignedto)
	require.Equal(t, bug1.Assignedbyprofile, bug2.Assignedbyprofile)
	require.Equal(t, bug1.Assignedbyteam, bug2.Assignedbyteam)
	require.Equal(t, bug1.Completed, bug2.Completed)
	require.WithinDuration(t, bug1.Createdat, bug2.Createdat, time.Second)
	require.WithinDuration(t, bug1.Updatedat, bug2.Updatedat, time.Second)
}

func TestUpdateBug(t *testing.T) {
	bug1, _, _ := createRandomBug(t)

	arg := UpdateBugParams{
		Name:        util.RandomString(10),
		Description: util.RandomString(100),
		Status:      BugstatusPENDING,
		Priority:    BugpriorityLOW,
		ID:          bug1.ID,
	}

	bug2, err := testQueries.UpdateBug(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bug2)
	require.Equal(t, arg.ID, bug2.ID)
	require.Equal(t, arg.Name, bug2.Name)
	require.Equal(t, arg.Description, bug2.Description)
	require.Equal(t, arg.Status, bug2.Status)
	require.Equal(t, arg.Priority, bug2.Priority)
	require.WithinDuration(t, bug1.Createdat, bug2.Createdat, time.Second)
}

func TestDeleteBug(t *testing.T) {
	bug1, _, _ := createRandomBug(t)

	arg := bug1.ID

	bug2, err := testQueries.DeleteBug(context.Background(), arg)
	require.NoError(t, err)
	require.NoError(t, err)
	require.NotEmpty(t, bug2)
	require.Equal(t, bug1.ID, bug2.ID)
	require.Equal(t, bug1.Name, bug2.Name)
	require.Equal(t, bug1.Description, bug2.Description)
	require.Equal(t, bug1.Status, bug2.Status)
	require.Equal(t, bug1.Priority, bug2.Priority)
	require.Equal(t, bug1.Assignedto, bug2.Assignedto)
	require.WithinDuration(t, bug1.Createdat, bug2.Createdat, time.Second)
	require.WithinDuration(t, bug1.Updatedat, bug2.Updatedat, time.Second)
}

func TestCloseBug(t *testing.T) {
	bug1, _, _ := createRandomBug(t)
	profile := createRandomProfile(t)

	arg := CloseBugParams{
		Closedby: pgtype.UUID{Bytes: profile.ID, Valid: true},
		Remarks:  pgtype.Text{String: util.RandomString(100), Valid: true},
		ID:       bug1.ID,
	}

	bug2, err := testQueries.CloseBug(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bug2)
	require.Equal(t, arg.Closedby, bug2.Closedby)
	require.Equal(t, arg.Remarks, bug2.Remarks)
}

func TestGetActiveBugsByAssignedProfile(t *testing.T) {
	bug1, _, member := createRandomBug(t)

	arg := member.Profile

	bugs, err := testQueries.GetActiveBugsByAssignedProfile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bugs)
	require.Contains(t, bugs, bug1)
}

func TestGetBugsByAssignedTeam(t *testing.T) {
	bug1, team, _ := createRandomBug(t)

	arg := team.ID

	bugs, err := testQueries.GetBugsByAssignedTeam(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bugs)
	require.Contains(t, bugs, bug1)
}

func TestGetBugsByAsigneeTeam(t *testing.T) {
	bug1, team, _ := createRandomBug(t)

	arg := team.ID

	bugs, err := testQueries.GetBugsByAssigneeTeam(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bugs)
	require.Contains(t, bugs, bug1)
}
