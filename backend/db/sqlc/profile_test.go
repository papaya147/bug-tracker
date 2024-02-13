package db

import (
	"context"
	"testing"
	"time"

	"github.com/papaya147/buggy/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomProfile(t *testing.T) Profile {
	arg := CreateProfileParams{
		ID:       util.RandomUuid(),
		Tokenid:  util.RandomUuid(),
		Name:     util.RandomString(10),
		Email:    util.RandomString(10),
		Password: util.RandomString(10),
	}

	profile, err := testQueries.CreateProfile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, profile)

	require.Equal(t, arg.ID, profile.ID)
	require.Equal(t, arg.Tokenid, profile.Tokenid)
	require.Equal(t, arg.Name, profile.Name)
	require.Equal(t, arg.Email, profile.Email)
	require.Equal(t, arg.Password, profile.Password)

	require.NotZero(t, profile.Createdat)
	require.NotZero(t, profile.Updatedat)

	return profile
}

func TestCreateProfile(t *testing.T) {
	createRandomProfile(t)
}

func TestGetProfile(t *testing.T) {
	profile1 := createRandomProfile(t)
	profile2, err := testQueries.GetProfile(context.Background(), profile1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, profile2)

	require.Equal(t, profile1.ID, profile2.ID)
	require.Equal(t, profile1.Tokenid, profile2.Tokenid)
	require.Equal(t, profile1.Name, profile2.Name)
	require.Equal(t, profile1.Email, profile2.Email)
	require.Equal(t, profile1.Password, profile2.Password)
	require.Equal(t, profile1.Createdat, profile2.Createdat)
	require.Equal(t, profile1.Updatedat, profile2.Updatedat)
}

func TestUpdatePassword(t *testing.T) {
	profile1 := createRandomProfile(t)

	arg := UpdatePasswordParams{
		Password: util.RandomString(10),
		ID:       profile1.ID,
	}

	profile2, err := testQueries.UpdatePassword(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, profile2)

	require.Equal(t, profile1.ID, profile2.ID)
	require.Equal(t, profile1.Tokenid, profile2.Tokenid)
	require.Equal(t, profile1.Name, profile2.Name)
	require.Equal(t, profile1.Email, profile2.Email)
	require.Equal(t, arg.Password, profile2.Password)
	require.Equal(t, profile1.Createdat, profile2.Createdat)
	require.WithinDuration(t, profile1.Updatedat, profile2.Updatedat, time.Second)
}

func TestUpdateProfile(t *testing.T) {
	profile1 := createRandomProfile(t)

	arg := UpdateProfileParams{
		Name: util.RandomString(10),
		ID:   profile1.ID,
	}

	profile2, err := testQueries.UpdateProfile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, profile2)

	require.Equal(t, profile1.ID, profile2.ID)
	require.Equal(t, profile1.Tokenid, profile2.Tokenid)
	require.Equal(t, arg.Name, profile2.Name)
	require.Equal(t, profile1.Email, profile2.Email)
	require.Equal(t, profile1.Password, profile2.Password)
	require.Equal(t, profile1.Createdat, profile2.Createdat)
	require.WithinDuration(t, profile1.Updatedat, profile2.Updatedat, time.Second)
}

func TestVerifyProfile(t *testing.T) {
	profile1 := createRandomProfile(t)

	profile2, err := testQueries.VerifyProfile(context.Background(), profile1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, profile2)

	require.Equal(t, profile1.ID, profile2.ID)
	require.Equal(t, profile1.Tokenid, profile2.Tokenid)
	require.Equal(t, profile1.Name, profile2.Name)
	require.Equal(t, profile1.Email, profile2.Email)
	require.Equal(t, profile1.Password, profile2.Password)
	require.Equal(t, profile1.Createdat, profile2.Createdat)
	require.WithinDuration(t, profile1.Updatedat, profile2.Updatedat, time.Second)

	require.NotEmpty(t, profile2.Verified)
}

func TestUpdateTokenId(t *testing.T) {
	profile1 := createRandomProfile(t)

	arg := UpdateTokenIdParams{
		Tokenid: util.RandomUuid(),
		ID:      profile1.ID,
	}

	profile2, err := testQueries.UpdateTokenId(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, profile2)

	require.Equal(t, profile1.ID, profile2.ID)
	require.Equal(t, arg.Tokenid, profile2.Tokenid)
	require.Equal(t, profile1.Name, profile2.Name)
	require.Equal(t, profile1.Email, profile2.Email)
	require.Equal(t, profile1.Password, profile2.Password)
	require.Equal(t, profile1.Createdat, profile2.Createdat)
	require.WithinDuration(t, profile1.Updatedat, profile2.Updatedat, time.Second)
}

func TestGetProfileByEmail(t *testing.T) {
	profile1 := createRandomProfile(t)

	profile2, err := testQueries.GetProfileByEmail(context.Background(), profile1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, profile2)

	require.Equal(t, profile1.ID, profile2.ID)
	require.Equal(t, profile2.Tokenid, profile2.Tokenid)
	require.Equal(t, profile1.Name, profile2.Name)
	require.Equal(t, profile1.Email, profile2.Email)
	require.Equal(t, profile1.Password, profile2.Password)
	require.Equal(t, profile1.Createdat, profile2.Createdat)
	require.WithinDuration(t, profile1.Updatedat, profile2.Updatedat, time.Second)

}

func TestUpdateTokenIdByEmail(t *testing.T) {
	profile1 := createRandomProfile(t)

	testQueries.VerifyProfile(context.Background(), profile1.ID)

	arg := UpdateTokenIdByEmailParams{
		Tokenid: util.RandomUuid(),
		Email:   profile1.Email,
	}

	profile2, err := testQueries.UpdateTokenIdByEmail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, profile2)

	require.Equal(t, profile1.ID, profile2.ID)
	require.Equal(t, profile2.Tokenid, profile2.Tokenid)
	require.Equal(t, profile1.Name, profile2.Name)
	require.Equal(t, profile1.Email, profile2.Email)
	require.Equal(t, profile1.Password, profile2.Password)
	require.Equal(t, profile1.Createdat, profile2.Createdat)
	require.WithinDuration(t, profile1.Updatedat, profile2.Updatedat, time.Second)
}
