package postgres

import (
	"context"
	"culand-internship/internal/core/postgres"
	"culand-internship/internal/internship/app"
	"github.com/jackc/pgx/v5"
)

type TxRunner struct {
	coreTx postgres.TxManager
}

func NewTxRunner(coreTx postgres.TxManager) *TxRunner {
	return &TxRunner{coreTx: coreTx}
}

func (r *TxRunner) WithinTx(ctx context.Context, fn func(context.Context, app.Repository) error) error {
	return r.coreTx.WithinTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		repo := NewRepository(tx)
		return fn(ctx, repo)
	})
}
