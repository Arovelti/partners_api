package handlers

import (
	"net/http"

	"partners_api/pkg/models"

	"github.com/gin-gonic/gin"
)

type PartnerHandler struct {
	partnerRepo models.PartnerUsecase
}

func (h *PartnerHandler) CreatePartnerHandler(c *gin.Context) {
	var partner models.Partner

	if err := c.BindJSON(&partner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.partnerRepo.CreatePartner(&partner); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, partner)
}

func (h *PartnerHandler) SetPartnerStatusHandler(c *gin.Context) {
	partnerID := c.Param("partner_id")
	if partnerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid partner ID"})
		return
	}

	var status models.Status

	if err := c.BindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.partnerRepo.SetPartnerStatus(partnerID, status.Active, status.ExpiryDate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *PartnerHandler) GetPartnerByIDHandler(c *gin.Context) {
	partnerID := c.Param("partner_id")
	if partnerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid partner ID"})
		return
	}

	partner, err := h.partnerRepo.GetPartnerByID(partnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, partner)
}

func (h *PartnerHandler) GetAllPartners(c *gin.Context) {
	partners, err := h.partnerRepo.GetAllPartners()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, partners)
}
