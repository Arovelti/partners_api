package db

import (
	"database/sql"
	"errors"
	"fmt"

	"partners_api/pkg/models"
)

var (
	ErrPartnerNotFound         = errors.New("partner with such id is not exist")
	ErrPartnerStatusIsInactive = errors.New("this partner is inactive now")
)

type PartnerPostgresRepo struct {
	db *sql.DB
}

func NewPartnerRepository(db *sql.DB) models.PartnerRepository {
	return &PartnerPostgresRepo{
		db: db,
	}
}

func (p *PartnerPostgresRepo) GetPartnerByID(id string) (*models.Partner, error) {
	query := `SELECT id, name, created_at, updated_at FROM partners WHERE id=$1`

	row := p.db.QueryRow(query, id)
	partner := &models.Partner{}
	err := row.Scan(&partner.ID, &partner.Name, &partner.CreatedAt, &partner.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrPartnerNotFound
		}
		return nil, err
	}

	if !partner.Status.Active {
		return nil, ErrPartnerStatusIsInactive
	}

	return partner, nil
}

func (p *PartnerPostgresRepo) GetAllPartners() ([]*models.Partner, error) {
	partners := []*models.Partner{}
	rows, err := p.db.Query("SELECT id, name, status, created_at, expires_at FROM partners")
	if err != nil {
		return nil, fmt.Errorf("failed to query partners: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		partner := &models.Partner{}
		err = rows.Scan(&partner.ID, &partner.Name, &partner.Status, &partner.CreatedAt, &partner.ExpiresAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan partner: %v", err)
		}

		if !partner.Status.Active {
			continue // if our partner is not active, we don't include him in our slice
		}

		partners = append(partners, partner)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over partners: %v", err)
	}

	return partners, nil
}

func (p *PartnerPostgresRepo) CreatePartner(partner *models.Partner) error {
	query := `
			INSERT INTO partners (id, name, email, created_at)
			VALUES ($1, $2, $3, $4)
		`

	_, err := p.db.Exec(query, partner.ID, partner.Name, partner.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (p *PartnerPostgresRepo) SetPartnerStatus(partner *models.Partner) error {
	query := `
        UPDATE partners SET name=$1, email=$2, updated_at=$3 WHERE id=$4
    `

	_, err := p.db.Exec(query, partner.Name, partner.UpdatedAt, partner.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *PartnerPostgresRepo) DeletePartner(id string) error {
	query := `DELETE FROM partners WHERE id=$1`

	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
