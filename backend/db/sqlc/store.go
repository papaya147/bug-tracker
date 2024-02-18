package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
	CreateOrganisationTransferTx(ctx context.Context, arg CreateOrganisationTransferTxParams) (Organisationtransfer, error)
	GetOrganisationTransfersTx(ctx context.Context, arg uuid.UUID) (GetOrganisationTransfersTxResponse, error)
	CompleteOrganisationTransferTx(ctx context.Context, arg CompleteOrganisationTransferTxParams) (Organisation, error)
}

type sqlStore struct {
	db *pgxpool.Pool
	*Queries
}

func NewStore(db *pgxpool.Pool) Store {
	return &sqlStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *sqlStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit(ctx)
}
