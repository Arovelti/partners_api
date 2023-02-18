package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"partners_api/pkg/models"
)

func GetPurchase(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var purchase models.Purchase

	err := json.NewDecoder(r.Body).Decode(&purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = sendPurchase(purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func sendPurchase(purchase models.Purchase) error {
	url := "http://example-endpoint"

	reqBody, err := json.Marshal(purchase)
	if err != nil {
		return err
	}

	_, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	return nil
}
