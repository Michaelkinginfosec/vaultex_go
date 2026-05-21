package repository

import (
	"context"
	"fmt"
	"vaultex/internal/model"
	"vaultex/pkg/util"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WalletRepository interface {
	Create(ctx context.Context, wallet *model.Wallet) error
	FindByOrganizationIDAndExternalUserID(ctx context.Context, organizationID, externalUserID string) (*model.Wallet, error)
}

type walletRepo struct {
	db *pgxpool.Pool
}

func NewWalletRepository(db *pgxpool.Pool) WalletRepository {
	return &walletRepo{db: db}
}

func (r *walletRepo) Create(ctx context.Context, wallet *model.Wallet) error {
	query := `INSERT INTO wallets (organization_id, external_user_id, external_user_email, currency, metadata) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(
		ctx,
		query,
		wallet.OrganizationID,
		wallet.ExternalUserID,
		wallet.ExternalUserEmail,
		wallet.Currency,
		wallet.MetaData,
	).Scan(
		&wallet.ID,
	)

	if err != nil {
		fmt.Printf("Error inserting wallet: %v\n", err)
		return util.InternalServerError("failed to create wallet")
	}

	return nil

}

func (r *walletRepo) FindByOrganizationIDAndExternalUserID(ctx context.Context, organizationID, externalUserID string) (*model.Wallet, error) {
	query := `SELECT id, organization_id, external_user_id, external_user_email, currency, balance, ledger_balance, metadata FROM wallets WHERE organization_id = $1 AND external_user_id = $2`
	wallet := &model.Wallet{}
	err := r.db.QueryRow(ctx, query, organizationID, externalUserID).Scan(&wallet.ID, &wallet.OrganizationID, &wallet.ExternalUserID, &wallet.ExternalUserEmail, &wallet.Currency, &wallet.Balance, &wallet.LedgerBalance, &wallet.MetaData)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, util.ErrNotFound
		}
		return nil, util.InternalServerError("failed to find wallet")
	}
	return wallet, nil
}
