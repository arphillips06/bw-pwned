package bitwarden

import (
	"bw-hibp-check/helper"
	"bw-hibp-check/models"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetStatus() {
	resp, err := http.Get("http://localhost:8087/status")
	if err != nil {
		log.Fatalf("Failed to get status: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	fmt.Println("Response:", string(body))
}

func UnlockVault(password string) (*models.UnlockResponse, error) {
	var resp models.UnlockResponse
	err := helper.DoRequest("POST", "http://localhost:8087/unlock",
		models.UnlockRequest{Password: password}, &resp)
	if err != nil {
		return nil, err
	}
	log.Printf("Unlocked: %v | Message: %s", resp.Success, resp.Data.Title)
	return &resp, nil
}

func GetItem(id string) (*models.BitwardenItemResponse, error) {
	var resp models.BitwardenItemResponse
	url := fmt.Sprintf("http://localhost:8087/object/item/%s", id)
	if err := helper.DoRequest("GET", url, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func ListAllItems() (*models.BitwardenItemsListResponse, error) {
	var resp models.BitwardenItemsListResponse
	err := helper.DoRequest("GET", "http://localhost:8087/object/items", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
