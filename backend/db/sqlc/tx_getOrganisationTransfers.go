package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/papaya147/buggy/backend/util"
)

type IncomingTransfer struct {
	Id                      uuid.UUID `json:"id"`
	OrganisationName        string    `json:"organisation_name"`
	OrganisationDescription string    `json:"organisation_description"`
	FromProfileName         string    `json:"from_profile_name"`
	CreatedAt               int64     `json:"created_at"`
}

type OutgoingTransfer struct {
	Id                      uuid.UUID `json:"id"`
	OrganisationName        string    `json:"organisation_name"`
	OrganisationDescription string    `json:"organisation_description"`
	ToProfileName           string    `json:"to_profile_name"`
	CreatedAt               int64     `json:"created_at"`
}

type GetOrganisationTransfersTxResponse struct {
	IncomingTransfers []IncomingTransfer `json:"incoming_transfers"`
	OutgoingTransfer  *OutgoingTransfer  `json:"outgoing_transfer"`
}

func (store *sqlStore) GetOrganisationTransfersTx(ctx context.Context, arg uuid.UUID) (GetOrganisationTransfersTxResponse, error) {
	iTransfers := []IncomingTransfer{}
	oTransfer := &OutgoingTransfer{}

	err := store.execTx(ctx, func(q *Queries) error {
		incomingTransfers, err := q.GetIncomingOrganisationTransfers(ctx, arg)
		if err != nil && err != pgx.ErrNoRows {
			return util.ErrDatabase
		}

		for _, transfer := range incomingTransfers {
			iTransfers = append(iTransfers, IncomingTransfer{
				Id:                      transfer.ID,
				OrganisationName:        transfer.Orgname,
				OrganisationDescription: transfer.Description,
				FromProfileName:         transfer.Profilename,
				CreatedAt:               transfer.Createdat.Unix(),
			})
		}

		outgoingTransfers, err := q.GetOutgoingOrganisationTransfers(ctx, arg)
		if err != nil && err != pgx.ErrNoRows {
			return util.ErrDatabase
		}

		if len(outgoingTransfers) == 0 {
			oTransfer = nil
			return nil
		}

		profile, err := q.GetProfile(ctx, outgoingTransfers[0].Toprofile)
		if err != nil {
			return util.ErrDatabase
		}

		org, err := q.GetOrganisationByOwner(ctx, arg)
		if err != nil && err != pgx.ErrNoRows {
			return util.ErrDatabase
		}

		oTransfer = &OutgoingTransfer{
			Id:                      outgoingTransfers[0].ID,
			OrganisationName:        org.Name,
			OrganisationDescription: org.Description,
			ToProfileName:           profile.Name,
			CreatedAt:               outgoingTransfers[0].Createdat.Unix(),
		}

		return nil
	})

	return GetOrganisationTransfersTxResponse{
		IncomingTransfers: iTransfers,
		OutgoingTransfer:  oTransfer,
	}, err
}
