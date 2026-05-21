package handlers

import (
	"fmt"
	"strconv"
	"time"
	"vaultex/internal/model"
	"vaultex/internal/service"
	"vaultex/internal/shared"
	"vaultex/pkg/dto"
	"vaultex/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type WalletHandler struct {
	service service.WalletService
}

func NewWalletHandler(s service.WalletService) *WalletHandler {
	return &WalletHandler{service: s}
}

// CreateWallet godoc
// @Summary Create user wallet
// @Description Create a wallet for an external user under the authenticated organization
// @Tags Wallets
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param x-api-key header string true "API Key"
// @Param x-signature header string true "HMAC Signature"
// @Param x-timestamp header string true "Unix Timestamp"
// @Param request body dto.CreateWalletRequest true "Create wallet request"
// @Success 201 {object} shared.APIResponse
// @Failure 400 {object} shared.APIResponse
// @Failure 401 {object} shared.APIResponse
// @Failure 409 {object} shared.APIResponse
// @Failure 500 {object} shared.APIResponse
// @Router /wallets [post]
func (h *WalletHandler) CreateWallet(c *gin.Context) {
	var req dto.CreateWalletRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			shared.BadRequest(c, "Validation failed", util.FormatValidationError(err))
		} else {
			shared.BadRequest(c, "Malformed request body", err.Error())
		}
		return
	}

	orgID := c.GetString("org_id")
	ctx := c.Request.Context()

	walletWithAccount, err := h.service.CreateWallet(ctx, orgID, &req)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	res := mapWalletWithAccountResponse(walletWithAccount)

	shared.Created(c, "Wallet created successfully", res)
}

// GetWalletByExternalUserID godoc
// @Summary Get wallet by external user ID
// @Description Retrieve a wallet using the authenticated organization and external user ID
// @Tags Wallets
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param x-api-key header string true "API Key"
// @Param x-signature header string true "HMAC Signature"
// @Param x-timestamp header string true "Unix Timestamp"
// @Param external_user_id path string true "External User ID"
// @Success 200 {object} shared.APIResponse
// @Failure 400 {object} shared.APIResponse
// @Failure 401 {object} shared.APIResponse
// @Failure 404 {object} shared.APIResponse
// @Failure 500 {object} shared.APIResponse
// @Router /wallets/{external_user_id} [get]
func (h *WalletHandler) GetWalletByExternalUserID(c *gin.Context) {
	organizationID := c.GetString("org_id")
	externalUserID := c.Param("external_user_id")

	ctx := c.Request.Context()

	walletWithAccount, err := h.service.GetWalletByOrganizationIDAndExternalUserID(
		ctx,
		organizationID,
		externalUserID,
	)

	if err != nil {
		shared.HandleError(c, err)
		return
	}

	res := mapWalletWithAccountResponse(walletWithAccount)

	shared.OK(c, "Wallet retrieved successfully", res)
}
func mapWalletWithAccountResponse(data *model.WalletWithAccount) *dto.WalletWithAccountResponse {
	return &dto.WalletWithAccountResponse{
		Wallet: dto.WalletResponse{
			ID:                data.Wallet.ID,
			OrganizationID:    data.Wallet.OrganizationID,
			ExternalUserID:    data.Wallet.ExternalUserID,
			ExternalUserEmail: data.Wallet.ExternalUserEmail,
			Currency:          data.Wallet.Currency,
			MetaData:          data.Wallet.MetaData,
			IsActive:          data.Wallet.IsActive,
			CreatedAt:         data.Wallet.CreatedAt.Format(time.RFC3339),
		},
		Account: dto.AccountResponse{
			ID:               data.Account.ID,
			WalletID:         data.Account.WalletID,
			AccountNumber:    data.Account.AccountNumber,
			AccountName:      data.Account.AccountName,
			Currency:         data.Account.Currency,
			LedgerBalance:    data.Account.LedgerBalance,
			LockedBalance:    data.Account.LockedBalance,
			AvailableBalance: calculateAvailableBalance(data.Account.LedgerBalance, data.Account.LockedBalance),
			IsActive:         data.Account.IsActive,
			CreatedAt:        data.Account.CreatedAt.Format(time.RFC3339),
		},
	}
}

func calculateAvailableBalance(ledgerBalance, lockedBalance string) string {
	ledger, _ := strconv.ParseFloat(ledgerBalance, 64)
	locked, _ := strconv.ParseFloat(lockedBalance, 64)

	return fmt.Sprintf("%.2f", ledger-locked)
}
