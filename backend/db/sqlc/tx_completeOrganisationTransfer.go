package db

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/util"
)

type CompleteOrganisationTransferTxParams struct {
	TransferId uuid.UUID
	ToProfile  uuid.UUID
}

func (store *sqlStore) CompleteOrganisationTransferTx(ctx context.Context, arg CompleteOrganisationTransferTxParams) (Organisation, error) {
	var transferredOrg Organisation

	err := store.execTx(ctx, func(q *Queries) error {
		// * check if profile already has organisation
		profileOrg, err := q.GetOrganisation(ctx, arg.ToProfile)
		if err != nil && err != pgx.ErrNoRows {
			return util.ErrDatabase
		}

		if profileOrg.ID != uuid.Nil {
			return util.ErrEntityExists
		}

		transfer, err := q.CompleteOrganisationTransfer(ctx, CompleteOrganisationTransferParams{
			ID:        arg.TransferId,
			Toprofile: arg.ToProfile,
		})
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return util.ErrEntityDoesNotExist
			}
			return util.ErrDatabase
		}

		transferredOrg, err = q.UpdateOrganisationOwner(ctx, UpdateOrganisationOwnerParams{
			Owner: arg.ToProfile,
			ID:    transfer.Organisation,
		})
		if err != nil {
			return util.ErrDatabase
		}

		return nil
	})

	return transferredOrg, err
}
