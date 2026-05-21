package service

import (
	"context"
	"vaultex/internal/model"
	"vaultex/internal/repository"
	"vaultex/pkg/dto"
	"vaultex/pkg/util"
)

type WalletService interface {
	CreateWallet(ctx context.Context, organizationID string, req *dto.CreateWalletRequest) (*model.WalletWithAccount, error)
	GetWalletByOrganizationIDAndExternalUserID(ctx context.Context, organizationID, externalUserID string) (*model.WalletWithAccount, error)
}

type walletService struct {
	WalletRepo repository.WalletRepository
}

func NewWalletService(ws repository.WalletRepository) WalletService {
	return &walletService{WalletRepo: ws}

}

func (s *walletService) CreateWallet(ctx context.Context, organizationID string, req *dto.CreateWalletRequest) (*model.WalletWithAccount, error) {
	existingWallet, err := s.WalletRepo.FindByOrganizationIDAndExternalUserID(ctx, organizationID, req.ExternalUserID)
	if err != nil && err != util.ErrNotFound {
		return nil, util.InternalServerError("failed to check existing wallet")
	}
	if existingWallet != nil {
		return nil, util.ConflictError("wallet already exists for this user")
	}

	wallet := &model.Wallet{
		OrganizationID:    organizationID,
		ExternalUserID:    req.ExternalUserID,
		ExternalUserEmail: req.ExternalUserEmail,
		Currency:          req.Currency,
		MetaData:          req.MetaData,
	}

	account := &model.Account{
		OrganizationID: organizationID,
		AccountNumber:  req.AccountNumber,
		AccountName:    req.AccountName,
		Currency:       req.Currency,
	}

	err = s.WalletRepo.Create(ctx, wallet, account)
	if err != nil {
		return nil, err
	}

	return &model.WalletWithAccount{
		Wallet:  wallet,
		Account: account,
	}, nil
}

func (s *walletService) GetWalletByOrganizationIDAndExternalUserID(ctx context.Context, organizationID, externalUserID string) (*model.WalletWithAccount, error) {
	wallet, err := s.WalletRepo.FindByOrganizationIDAndExternalUserID(ctx, organizationID, externalUserID)

	if err != nil {
		if err == util.ErrNotFound {
			return nil, util.NotFoundError("wallet not found")
		}

		return nil, err
	}

	return wallet, nil
}
