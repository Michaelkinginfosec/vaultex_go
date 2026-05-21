package service

import (
	"context"
	"fmt"
	"strings"
	"vaultex/internal/model"
	"vaultex/internal/repository"
	"vaultex/pkg/util"
)

type Service interface {
	CreateOrganization(ctx context.Context, organizationname string, registeredname string, phonenumber string, email string, websiteurl string, password string) (*model.Organization, error)
	FindOrganizationByEmail(ctx context.Context, email string) (*model.Organization, error)
	Login(ctx context.Context, email string, password string) (*model.Organization, error)
}

type service struct {
	OrgRepo repository.OrganizationRepository
}

func NewService(ur repository.OrganizationRepository) Service {
	return &service{OrgRepo: ur}
}

func (s *service) CreateOrganization(ctx context.Context, organizationname string, registeredname string, phonenumber string, email string, websiteurl string, password string) (*model.Organization, error) {

	existingOrg, _ := s.OrgRepo.FindByEmail(ctx, email)
	if existingOrg != nil {
		return nil, util.ConflictError("organization with this email already exists")
	}

	apiKey, err := util.GenerateAPIKey()
	if err != nil {
		return nil, util.InternalServerError("failed to generate api key")
	}

	apiSecret, err := util.GenerateAPISecretKey()
	if err != nil {
		return nil, util.InternalServerError("failed to generate api secret")
	}

	encryptedSecret, err := util.EncryptSecret(apiSecret)
	if err != nil {
		return nil, util.InternalServerError("failed to encrypt api secret")
	}

	hashPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, util.InternalServerError("failed to hash password")
	}
	organization := &model.Organization{
		OrganizationName: organizationname,
		RegisteredName:   registeredname,
		PhoneNumber:      phonenumber,
		WebsiteURL:       &websiteurl,
		APIKey:           apiKey,
		APISecret:        encryptedSecret,
		Email:            email,
		Password:         hashPassword,
	}

	err = s.OrgRepo.Create(ctx, organization)
	if err != nil {
		return nil, util.InternalServerError("failed to create organization")
	}
	organization.APISecret = apiSecret
	return organization, nil
}

func (s *service) FindOrganizationByEmail(ctx context.Context, email string) (*model.Organization, error) {
	org, err := s.OrgRepo.FindByEmail(ctx, email)

	if err != nil {
		if err == util.ErrNotFound {
			return nil, util.NotFoundError("organization not found")
		}
		return nil, util.InternalServerError(err.Error())
	}
	return org, nil
}
func (s *service) Login(ctx context.Context, email string, password string) (*model.Organization, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	org, err := s.OrgRepo.Login(ctx, email)
	if err != nil {
		fmt.Print(err)
		return nil, util.UnauthorizedError("invalid credentials")
	}

	// fmt.Println("ORG ID:", org.ID)
	// fmt.Printf("EMAIL: %q\n", org.Email)
	// fmt.Printf("PLAIN PASSWORD: %q\n", password)
	// fmt.Printf("HASH FROM DB: %q\n", org.Password)
	// fmt.Println("HASH LENGTH:", len(org.Password))
	// fmt.Println("COMPARE RESULT:", util.ComparePassword(org.Password, password))
	if !util.ComparePassword(org.Password, password) {
		fmt.Printf("Password mismatch for email: %s\n", email)
		return nil, util.UnauthorizedError("invalid credentials")
	}
	return org, nil
}
