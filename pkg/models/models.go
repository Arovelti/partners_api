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

// swagger:parameters createPurchase
type Purchase struct {
	// Token ID for the purchase
	//
	// in: body
	// required: true
	TokenID string `json:"token_id"`

	// Contract ID for the purchase
	//
	// in: body
	// required: true
	ContractID string `json:"contract_id"`

	// Partner name for the purchase
	//
	// in: body
	// required: true
	PartnerID int `json:"partner_id"`

	// Amount for the purchase
	//
	// in: body
	// required: true
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
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
