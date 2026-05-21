package handlers

import (
	"time"
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
	ctx := c.Request.Context()

	wallet, err := h.service.CreateWallet(ctx, &req)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	res := &dto.CreateWalletResponse{
		ID:                wallet.ID,
		OrganizationID:    wallet.OrganizationID,
		ExternalUserID:    wallet.ExternalUserID,
		ExternalUserEmail: wallet.ExternalUserEmail,
		Currency:          wallet.Currency,
		Balance:           wallet.Balance,
		LedgerBalance:     wallet.LedgerBalance,
		MetaData:          wallet.MetaData,
		IsActive:          wallet.IsActive,
		CreatedAt:         wallet.CreatedAt.Format(time.RFC3339),
	}
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
func (h *WalletHandler) GetWalletByOrganizationIDAndExternalUserID(c *gin.Context) {
	organizationID := c.GetString("org_id")
	externalUserID := c.Param("external_user_id")

	ctx := c.Request.Context()
	wallet, err := h.service.GetWalletByOrganizationIDAndExternalUserID(ctx, organizationID, externalUserID)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	res := &dto.CreateWalletResponse{
		ID:                wallet.ID,
		OrganizationID:    wallet.OrganizationID,
		ExternalUserID:    wallet.ExternalUserID,
		ExternalUserEmail: wallet.ExternalUserEmail,
		Currency:          wallet.Currency,
		Balance:           wallet.Balance,
		LedgerBalance:     wallet.LedgerBalance,
		MetaData:          wallet.MetaData,
		IsActive:          wallet.IsActive,
		CreatedAt:         wallet.CreatedAt.Format(time.RFC3339),
	}
	shared.OK(c, "Wallet retrieved successfully", res)
}
