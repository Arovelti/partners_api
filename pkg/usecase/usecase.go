package usecase

import (
	"time"

	"partners_api/pkg/models"
)

type PartnerUsecase struct {
	partnerRepo models.PartnerRepository
}

func NewPartnerUsecase(repo models.PartnerRepository) models.PartnerUsecase {
	return &PartnerUsecase{
		partnerRepo: repo,
	}
}

func (p *PartnerUsecase) GetPartnerByID(id string) (*models.Partner, error) {
	partner, err := p.partnerRepo.GetPartnerByID(id)
	if err != nil {
		return nil, err
	}

	return partner, nil
}

func (p *PartnerUsecase) GetAllPartners() ([]*models.Partner, error) {
	partners, err := p.partnerRepo.GetAllPartners()
	if err != nil {
		return nil, err
	}

	return partners, err
}

func (p *PartnerUsecase) CreatePartner(partner *models.Partner) error {
	err := p.partnerRepo.CreatePartner(partner)
	if err != nil {
		return err
	}

	return nil
}

func (p *PartnerUsecase) SetPartnerStatus(id string, active bool, expiryDate time.Time) error {
	partner, err := p.partnerRepo.GetPartnerByID(id)
	if err != nil {
		return err
	}

	partner.Status.Active = active
	partner.Status.ExpiryDate = expiryDate

	err = p.partnerRepo.SetPartnerStatus(partner)
	if err != nil {
		return err
	}

	return nil
}
