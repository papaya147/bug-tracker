package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/util"
)

type CompleteOrganisationTransferTxParams struct {
	TransferId uuid.UUID
	ToProfile  uuid.UUID
}

func (store *sqlStore) CompleteOrganisationTransferTx(ctx context.Context, arg CompleteOrganisationTransferTxParams) (Organisation, error) {
	var org Organisation

	err := store.execTx(ctx, func(q *Queries) error {
		transfer, err := q.CompleteOrganisationTransfer(ctx, CompleteOrganisationTransferParams{
			ID:        arg.TransferId,
			Toprofile: arg.ToProfile,
		})
		if err != nil {
			if err == pgx.ErrNoRows {
				return util.ErrEntityDoesNotExist
			}
			return util.ErrDatabase
		}

		org, err = q.UpdateOrganisationOwner(ctx, UpdateOrganisationOwnerParams{
			Owner: arg.ToProfile,
			ID:    transfer.Organisation,
		})
		if err != nil {
			return util.ErrDatabase
		}

		return nil
	})

	return org, err
}
