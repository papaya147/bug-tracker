package db

import (
	"context"

	"github.com/google/uuid"
)

type mockStore struct {
}

// CloseBug implements Store.
func (m *mockStore) CloseBug(ctx context.Context, arg CloseBugParams) (Bug, error) {
	return Bug{}, nil
}

// CreateBug implements Store.
func (m *mockStore) CreateBug(ctx context.Context, arg CreateBugParams) (Bug, error) {
	return Bug{}, nil
}

// DeleteBug implements Store.
func (m *mockStore) DeleteBug(ctx context.Context, id uuid.UUID) (Bug, error) {
	return Bug{}, nil
}

// GetActiveBugsByProfile implements Store.
func (m *mockStore) GetActiveBugsByProfile(ctx context.Context, profile uuid.UUID) ([]Bug, error) {
	return nil, nil
}

// GetBug implements Store.
func (m *mockStore) GetBug(ctx context.Context, id uuid.UUID) (Bug, error) {
	return Bug{}, nil
}

// GetBugsByAsigneeTeam implements Store.
func (m *mockStore) GetBugsByAsigneeTeam(ctx context.Context, team uuid.UUID) ([]Bug, error) {
	return nil, nil
}

// GetBugsByAssignedTeam implements Store.
func (m *mockStore) GetBugsByAssignedTeam(ctx context.Context, assignedto uuid.UUID) ([]Bug, error) {
	return nil, nil
}

// UpdateBug implements Store.
func (m *mockStore) UpdateBug(ctx context.Context, arg UpdateBugParams) (Bug, error) {
	return Bug{}, nil
}

// GetOrganisation implements Store.
func (m *mockStore) GetOrganisation(ctx context.Context, id uuid.UUID) (Organisation, error) {
	return Organisation{}, nil
}

// GetTeams implements Store.
func (m *mockStore) GetTeams(ctx context.Context, profile uuid.UUID) ([]GetTeamsRow, error) {
	return nil, nil
}

// GetAllTeamMembers implements Store.
func (*mockStore) GetAllTeamMembers(ctx context.Context, team uuid.UUID) ([]GetAllTeamMembersRow, error) {
	return nil, nil
}

// UpdateTeamMember implements Store.
func (*mockStore) UpdateTeamMember(ctx context.Context, arg UpdateTeamMemberParams) (Teammember, error) {
	return Teammember{}, nil
}

// GetTeamOrganisation implements Store.
func (*mockStore) GetTeamOrganisation(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	return uuid.Nil, nil
}

// GetTeamMember implements Store.
func (*mockStore) GetTeamMember(ctx context.Context, arg GetTeamMemberParams) (Teammember, error) {
	return Teammember{
		Admin: true,
	}, nil
}

// CreateTeamMember implements Store.
func (*mockStore) CreateTeamMember(ctx context.Context, arg CreateTeamMemberParams) (Teammember, error) {
	return Teammember{}, nil
}

// UpdateTeam implements Store.
func (*mockStore) UpdateTeam(ctx context.Context, arg UpdateTeamParams) (Team, error) {
	return Team{}, nil
}

// GetOrganisationTeams implements Store.
func (*mockStore) GetOrganisationTeams(ctx context.Context, organisation uuid.UUID) ([]Team, error) {
	return nil, nil
}

// CreateTeam implements Store.
func (*mockStore) CreateTeam(ctx context.Context, arg CreateTeamParams) (Team, error) {
	return Team{}, nil
}

// CompleteOrganisationTransferTx implements Store.
func (*mockStore) CompleteOrganisationTransferTx(ctx context.Context, arg CompleteOrganisationTransferTxParams) (Organisation, error) {
	return Organisation{}, nil
}

// UpdateOrganisationOwner implements Store.
func (*mockStore) UpdateOrganisationOwner(ctx context.Context, arg UpdateOrganisationOwnerParams) (Organisation, error) {
	return Organisation{}, nil
}

// CompleteOrganisationTransfer implements Store.
func (*mockStore) CompleteOrganisationTransfer(ctx context.Context, arg CompleteOrganisationTransferParams) (Organisationtransfer, error) {
	return Organisationtransfer{}, nil
}

// DeleteOrganisationTransfer implements Store.
func (*mockStore) DeleteOrganisationTransfer(ctx context.Context, arg uuid.UUID) (Organisationtransfer, error) {
	return Organisationtransfer{}, nil
}

// GetOrganisationTransfersTx implements Store.
func (*mockStore) GetOrganisationTransfersTx(ctx context.Context, arg uuid.UUID) (GetOrganisationTransfersTxResponse, error) {
	return GetOrganisationTransfersTxResponse{}, nil
}

// GetIncomingOrganisationTransfers implements Store.
func (*mockStore) GetIncomingOrganisationTransfers(ctx context.Context, toprofile uuid.UUID) ([]GetIncomingOrganisationTransfersRow, error) {
	return nil, nil
}

// GetOutgoingOrganisationTransfers implements Store.
func (*mockStore) GetOutgoingOrganisationTransfers(ctx context.Context, fromprofile uuid.UUID) ([]Organisationtransfer, error) {
	return nil, nil
}

// CreateOrganisationTransferTx implements Store.
func (*mockStore) CreateOrganisationTransferTx(ctx context.Context, arg CreateOrganisationTransferTxParams) (Organisationtransfer, error) {
	return Organisationtransfer{}, nil
}

// GetActiveOrganisationTransfer implements Store.
func (*mockStore) GetActiveOrganisationTransfer(ctx context.Context, organisation uuid.UUID) (Organisationtransfer, error) {
	return Organisationtransfer{}, nil
}

// CreateOrganisationTransfer implements Store.
func (*mockStore) CreateOrganisationTransfer(ctx context.Context, arg CreateOrganisationTransferParams) (Organisationtransfer, error) {
	return Organisationtransfer{}, nil
}

// UpdateOrganisation implements Store.
func (*mockStore) UpdateOrganisation(ctx context.Context, arg UpdateOrganisationParams) (Organisation, error) {
	return Organisation{}, nil
}

// GetOrganisation implements Store.
func (*mockStore) GetOrganisationByOwner(ctx context.Context, id uuid.UUID) (Organisation, error) {
	return Organisation{}, nil
}

// CreateOrganisation implements Store.
func (*mockStore) CreateOrganisation(ctx context.Context, arg CreateOrganisationParams) (Organisation, error) {
	return Organisation{}, nil
}

// CreateProfile implements Store.
func (*mockStore) CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error) {
	return Profile{}, nil
}

// GetProfile implements Store.
func (*mockStore) GetProfile(ctx context.Context, id uuid.UUID) (Profile, error) {
	return Profile{
		Password: "$2a$12$e.LqnwKjKFqzh8PVqz7r8..qtaFNBLtQbsYuEVrtObMwFlwNYLe3y",
	}, nil
}

// GetProfileByEmail implements Store.
func (*mockStore) GetProfileByEmail(ctx context.Context, email string) (Profile, error) {
	return Profile{}, nil
}

// UpdatePassword implements Store.
func (*mockStore) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (Profile, error) {
	return Profile{}, nil
}

// UpdateProfile implements Store.
func (*mockStore) UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error) {
	return Profile{}, nil
}

// UpdateTokenId implements Store.
func (*mockStore) UpdateTokenId(ctx context.Context, arg UpdateTokenIdParams) (Profile, error) {
	return Profile{}, nil
}

// UpdateTokenIdByEmail implements Store.
func (*mockStore) UpdateTokenIdByEmail(ctx context.Context, arg UpdateTokenIdByEmailParams) (Profile, error) {
	return Profile{
		Password: "$2a$12$e.LqnwKjKFqzh8PVqz7r8..qtaFNBLtQbsYuEVrtObMwFlwNYLe3y",
	}, nil
}

// VerifyProfile implements Store.
func (*mockStore) VerifyProfile(ctx context.Context, id uuid.UUID) (Profile, error) {
	return Profile{}, nil
}

func NewMockStore() Store {
	return &mockStore{}
}
