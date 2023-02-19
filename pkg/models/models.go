package models

import "time"

type Partner struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	ExpiresAt string `json:"expired,omitempty"`
	Status    Status `json:"status,omitempty"`
}

type Status struct {
	Active     bool      `json:"active"`
	ExpiryDate time.Time `json:"expiryDate"`
}

type Purchase struct {
	TokenID         int       `json:"token_id"`
	ContractAddress string    `json:"contract_address"`
	PartnerName     string    `json:"partner_name"`
	PurchaseAmount  float64   `json:"purchase_amount"`
	Timestamp       time.Time `json:"timestamp"`
}

// DB
type PartnerRepository interface {
	GetPartnerByID(id string) (*Partner, error)
	GetAllPartners() ([]*Partner, error)
	CreatePartner(partner *Partner) error
	SetPartnerStatus(partner *Partner) error
}

// Usecase
type PartnerUsecase interface {
	GetPartnerByID(id string) (*Partner, error)
	GetAllPartners() ([]*Partner, error)
	CreatePartner(partner *Partner) error
	SetPartnerStatus(id string, active bool, expiryDate time.Time) error
}
