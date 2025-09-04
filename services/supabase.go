package services

import (
	"backend/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type FoodItem struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Quantity   string `json:"quantity"`
	ExpiryDate string `json:"expiry_date"`
	UserID     string `json:"user_id"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func GetExpiringItems(cfg config.Config) ([]FoodItem, error) {
	url := fmt.Sprintf("%s/rest/v1/food_items?select=*&expiry_date=lte.%s",
		cfg.SupabaseUrl, time.Now().Add(48*time.Hour).Format("2006-01-02"))

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", cfg.SupabaseServiceKey)
	req.Header.Set("Authorization", "Bearer "+cfg.SupabaseServiceKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var items []FoodItem
	if err := json.Unmarshal(body, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func GetUserEmail(cfg config.Config, userId string) (string, error) {
	url := fmt.Sprintf("%s/auth/v1/admin/users/%s", cfg.SupabaseUrl, userId)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", cfg.SupabaseServiceKey)
	req.Header.Set("Authorization", "Bearer "+cfg.SupabaseServiceKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var user UserResponse
	if err := json.Unmarshal(body, &user); err != nil {
		return "", err
	}

	return user.Email, nil
}
