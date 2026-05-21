package service

import (
	"context"
	"fmt"
	"vaultex/internal/model"
	"vaultex/internal/repository"
	"vaultex/pkg/dto"
	"vaultex/pkg/util"
)

type WalletService interface {
	CreateWallet(ctx context.Context, req *dto.CreateWalletRequest) (*model.Wallet, error)
	GetWalletByOrganizationIDAndExternalUserID(ctx context.Context, organizationID, externalUserID string) (*model.Wallet, error)
}

type walletService struct {
	WalletRepo repository.WalletRepository
}

func NewWalletService(ws repository.WalletRepository) WalletService {
	return &walletService{WalletRepo: ws}

}

func (s *walletService) CreateWallet(ctx context.Context, req *dto.CreateWalletRequest) (*model.Wallet, error) {
	existingOrg, _ := s.WalletRepo.FindByOrganizationIDAndExternalUserID(ctx, req.OrganizationID, req.ExternalUserID)
	if existingOrg != nil {
		return nil, util.ConflictError("Wallet already exists for this user")
	}
	wallet := &model.Wallet{
		OrganizationID:    req.OrganizationID,
		ExternalUserID:    req.ExternalUserID,
		ExternalUserEmail: req.ExternalUserEmail,
		Currency:          req.Currency,
		MetaData:          req.MetaData,
	}

	err := s.WalletRepo.Create(ctx, wallet)
	if err != nil {
		fmt.Printf("Error creating wallet: %v\n", err)
		return nil, util.InternalServerError("failed to create wallet")
	}
	return wallet, nil
}

func (s *walletService) GetWalletByOrganizationIDAndExternalUserID(ctx context.Context, organizationID, externalUserID string) (*model.Wallet, error) {
	wallet, err := s.WalletRepo.FindByOrganizationIDAndExternalUserID(ctx, organizationID, externalUserID)
	if err != nil {
		return nil, util.InternalServerError("failed to find wallet")
	}

	userWallet := &model.Wallet{
		ID:                wallet.ID,
		OrganizationID:    wallet.OrganizationID,
		ExternalUserID:    wallet.ExternalUserID,
		ExternalUserEmail: wallet.ExternalUserEmail,
		Currency:          wallet.Currency,
		Balance:           wallet.Balance,
		LedgerBalance:     wallet.LedgerBalance,
		MetaData:          wallet.MetaData,
		IsActive:          wallet.IsActive,
		CreatedAt:         wallet.CreatedAt,
	}
	return userWallet, nil
}
