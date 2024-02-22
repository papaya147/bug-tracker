package db

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/util"
)

type CreateOrganisationTransferTxParams struct {
	FromProfile uuid.UUID
	ToEmail     string
}

func (store *sqlStore) CreateOrganisationTransferTx(ctx context.Context, arg CreateOrganisationTransferTxParams) (Organisationtransfer, error) {
	var transfer Organisationtransfer

	err := store.execTx(ctx, func(q *Queries) error {
		toProfile, err := q.GetProfileByEmail(ctx, arg.ToEmail)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return util.ErrProfileNotFound
			}
			return util.ErrDatabase
		}

		if toProfile.ID == arg.FromProfile {
			return util.ErrCannotTransferToSelf
		}

		org, err := q.GetOrganisation(ctx, arg.FromProfile)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return util.ErrEntityDoesNotExist
			}
			return util.ErrDatabase
		}

		activeTransfers, err := q.GetActiveOrganisationTransfer(ctx, org.ID)
		if err != nil && err != pgx.ErrNoRows {
			return util.ErrDatabase
		}

		// * checking if active transfer exists
		if activeTransfers.Organisation == org.ID {
			return util.ErrTransferAlreadyExists
		}

		tranferId, err := uuid.NewV7()
		if err != nil {
			return util.ErrInternal
		}

		transfer, err = q.CreateOrganisationTransfer(ctx, CreateOrganisationTransferParams{
			ID:           tranferId,
			Organisation: org.ID,
			Fromprofile:  arg.FromProfile,
			Toprofile:    toProfile.ID,
		})
		if err != nil {
			return util.ErrDatabase
		}

		return nil
	})

	return transfer, err
}
