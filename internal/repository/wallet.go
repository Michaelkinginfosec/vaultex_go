package repository

import (
	"context"
	"vaultex/internal/model"
	"vaultex/pkg/util"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WalletRepository interface {
	Create(ctx context.Context, wallet *model.Wallet, account *model.Account) error
	FindByOrganizationIDAndExternalUserID(ctx context.Context, organizationID, externalUserID string) (*model.WalletWithAccount, error)
}

type walletRepo struct {
	db *pgxpool.Pool
}

func NewWalletRepository(db *pgxpool.Pool) WalletRepository {
	return &walletRepo{db: db}
}

func (r *walletRepo) Create(ctx context.Context, wallet *model.Wallet, account *model.Account) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return util.InternalServerError("failed to start transaction")
	}
	defer tx.Rollback(ctx)

	walletQuery := `INSERT INTO wallets (organization_id, external_user_id, external_user_email, currency, metadata) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
	err = tx.QueryRow(
		ctx,
		walletQuery,
		wallet.OrganizationID,
		wallet.ExternalUserID,
		wallet.ExternalUserEmail,
		wallet.Currency,
		wallet.MetaData,
	).Scan(
		&wallet.ID,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)

	if err != nil {
		return util.InternalServerError("failed to create wallet")
	}
	accountQuery := `INSERT INTO accounts (wallet_id, organization_id, account_number,currency) VALUES ($1, $2, $3, $4) RETURNING id, ledger_balance, locked_balance, is_active, created_at, updated_at`

	account.WalletID = wallet.ID

	err = tx.QueryRow(
		ctx,
		accountQuery,
		account.WalletID,
		account.OrganizationID,
		account.AccountNumber,
		account.AccountName,
		account.Currency,
	).Scan(
		&account.ID,
		&account.LedgerBalance,
		&account.LockedBalance,
		&account.IsActive,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil {
		return util.InternalServerError("failed to create account")
	}
	if err := tx.Commit(ctx); err != nil {
		return util.InternalServerError("failed to commit transaction")
	}

	return nil

}

func (r *walletRepo) FindByOrganizationIDAndExternalUserID(ctx context.Context, organizationID, externalUserID string) (*model.WalletWithAccount, error) {
	query := `
	SELECT
		w.id,
		w.organization_id,
		w.external_user_id,
		w.external_user_email,
		w.currency,
		w.is_active,
		w.metadata,
		w.created_at,
		w.updated_at,

		a.id,
		a.wallet_id,
		a.organization_id,
		a.account_number,
		a.account_name,
		a.currency,
		a.ledger_balance,
		a.locked_balance,
		a.is_active,
		a.created_at,
		a.updated_at

	FROM wallets w
	INNER JOIN accounts a
		ON a.wallet_id = w.id

	WHERE w.organization_id = $1
	AND w.external_user_id = $2
`
	wallet := &model.Wallet{}
	account := &model.Account{}
	err := r.db.QueryRow(ctx, query, organizationID, externalUserID).Scan(
		&wallet.ID,
		&wallet.OrganizationID,
		&wallet.ExternalUserID,
		&wallet.ExternalUserEmail,
		&wallet.Currency,
		&wallet.IsActive,
		&wallet.MetaData,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,

		&account.ID,
		&account.WalletID,
		&account.OrganizationID,
		&account.AccountNumber,
		&account.AccountName,
		&account.Currency,
		&account.LedgerBalance,
		&account.LockedBalance,
		&account.IsActive,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, util.ErrNotFound
		}

		return nil, util.InternalServerError("failed to fetch wallet")
	}

	return &model.WalletWithAccount{Wallet: wallet, Account: account}, nil
}
