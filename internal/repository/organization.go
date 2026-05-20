package repository

import (
	"context"
	"vaultex/internal/model"
	"vaultex/pkg/util"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrganizationRepository interface {
	Create(ctx context.Context, organization *model.Organization) error
	FindByEmail(ctx context.Context, email string) (*model.Organization, error)
	Login(ctx context.Context, email string) (*model.Organization, error)
}

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) OrganizationRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, organization *model.Organization) error {
	query := `INSERT INTO organization (organization_name, registered_name, phone_number, website_url,  api_key, api_secret, email, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	return r.db.QueryRow(ctx, query, organization.OrganizationName, organization.RegisteredName, organization.PhoneNumber, organization.WebsiteURL, organization.APIKey, organization.APISecret, organization.Email, organization.Password).Scan(&organization.ID)
}

func (r *repo) FindByEmail(ctx context.Context, email string) (*model.Organization, error) {
	var organization model.Organization
	query := `SELECT id, organization_name, registered_name, phone_number, email, api_key FROM organization WHERE email = $1`
	err := r.db.QueryRow(ctx, query, email).Scan(&organization.ID, &organization.OrganizationName, &organization.RegisteredName, &organization.PhoneNumber, &organization.Email, &organization.APIKey)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, util.ErrNotFound
		}
		return nil, util.InternalServerError(err.Error())
	}
	return &organization, nil
}

func (r *repo) Login(ctx context.Context, email string) (*model.Organization, error) {
	var organization model.Organization
	query := `SELECT id, organization_name, registered_name, phone_number, email, api_key, password FROM organization WHERE email = $1`
	err := r.db.QueryRow(ctx, query, email).Scan(&organization.ID, &organization.OrganizationName, &organization.RegisteredName, &organization.PhoneNumber, &organization.Email, &organization.APIKey, &organization.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, util.ErrNotFound
		}
		return nil, util.InternalServerError(err.Error())
	}
	return &organization, nil
}
