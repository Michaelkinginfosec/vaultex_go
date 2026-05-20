package handlers

import (
	"net/http"
	"vaultex/internal/service"
	"vaultex/internal/shared"
	"vaultex/pkg/dto"
	"vaultex/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OrganizationHandler struct {
	service service.Service
}

func NewOrganizationHandler(s service.Service) *OrganizationHandler {
	return &OrganizationHandler{service: s}
}

// CreateOrganization godoc
// @Summary Create a new organization
// @Description Create a new organization with the provided name and email
// @Tags Auth
// @Accept json
// @Produce json
// @Param organization body dto.CreateOrganizationRequest true "Organization data"
// @Success 201 {object} dto.OrganizationSignupResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /organizations [post]
func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {
	var req dto.CreateOrganizationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			shared.BadRequest(c, "Validation failed", util.FormatValidationError(err))
		} else {
			shared.BadRequest(c, "Malformed request body", err.Error())
		}
		return
	}

	ctx := c.Request.Context()

	organization, err := h.service.CreateOrganization(
		ctx,
		req.OrganizationName,
		req.RegisteredName,
		req.PhoneNumber,
		req.Email,
		req.WebsiteURL,
		req.Password,
	)

	if err != nil {
		shared.HandleError(c, err)
		return
	}

	res := dto.OrganizationSignupResponse{
		ID:               organization.ID,
		OrganizationName: organization.OrganizationName,
		RegisteredName:   organization.RegisteredName,
		PhoneNumber:      organization.PhoneNumber,
		Email:            organization.Email,
		APIKey:           organization.APIKey,
		APISecret:        organization.APISecret,
	}

	shared.Created(c, "Organization created successfully", res)
}

// FindOrganizationByEmail godoc
// @Summary Get organization by email
// @Description Retrieve an organization by their email address
// @Tags organizations
// @Accept json
// @Produce json
// @Param email path string true "Organization Email"
// @Success 200 {object} dto.OrganizationLoginResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /organizations/email/{email} [get]
func (h *OrganizationHandler) FindOrganizationByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	organization, err := h.service.FindOrganizationByEmail(c, email)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	if organization == nil {
		shared.NotFound(c, "Organization not found", nil)
		return
	}
	res := dto.OrganizationLoginResponse{
		ID:               organization.ID,
		OrganizationName: organization.OrganizationName,
		RegisteredName:   organization.RegisteredName,
		PhoneNumber:      organization.PhoneNumber,
		Email:            organization.Email,
		APIKey:           organization.APIKey,
	}
	c.JSON(http.StatusOK, res)
}

// Login godoc
// @Summary Login an organization
// @Description Authenticate an organization using their email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.OrganizationLoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /organizations/login [post]
func (h *OrganizationHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			shared.BadRequest(c, "Validation failed", util.FormatValidationError(err))

		} else {
			shared.BadRequest(c, "Malformed request body", err.Error())
		}
		return
	}
	ctx := c.Request.Context()

	organization, err := h.service.Login(ctx, req.Email, req.Password)

	if err != nil {
		shared.Unauthorized(c, "Invalid email or password", nil)
		return
	}

	res := dto.OrganizationLoginResponse{
		ID:               organization.ID,
		OrganizationName: organization.OrganizationName,
		RegisteredName:   organization.RegisteredName,
		PhoneNumber:      organization.PhoneNumber,
		Email:            organization.Email,
		APIKey:           organization.APIKey,
	}

	shared.OK(c, "Organization logged in successfully", res)
}
