package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"partners_api/pkg/models"

	"github.com/gin-gonic/gin"
)

func GetPurchase(c *gin.Context) {
	var purchase models.Purchase
	if err := c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	reqData := struct {
		TokenID         int    `json:"token_id"`
		ContractAddress string `json:"contract_address"`
	}{
		TokenID:         purchase.TokenID,
		ContractAddress: purchase.ContractAddress,
	}

	reqDataJSON, err := json.Marshal(reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode request data"})
		return
	}

	url := "https://user121669781-n3gp3klg.wormhole.vk-apps.com/v1/token/updateScore"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqDataJSON))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	adminToken := "ET0yObYUj9"
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("admin_token", adminToken)

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
		return
	}
	defer resp.Body.Close()

	fmt.Println(purchase)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
