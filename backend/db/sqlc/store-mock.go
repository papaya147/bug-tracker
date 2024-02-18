package db

import (
	"context"

	"github.com/google/uuid"
)

type mockStore struct {
}

// DeleteOrganisationTransfer implements Store.
func (*mockStore) DeleteOrganisationTransfer(ctx context.Context, arg DeleteOrganisationTransferParams) (Organisationtransfer, error) {
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
func (*mockStore) GetOrganisation(ctx context.Context, id uuid.UUID) (Organisation, error) {
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
