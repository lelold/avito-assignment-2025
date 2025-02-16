package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestSendCoins(t *testing.T) {
	senderToken := authenticateUser("sender", "password")
	receiverToken := authenticateUser("receiver", "password")

	if senderToken == "" || receiverToken == "" {
		t.Fatal("Ошибка аутентификации пользователей")
	}

	sendCoinsReq := map[string]interface{}{
		"toUser": "receiver",
		"amount": 50,
	}
	body, _ := json.Marshal(sendCoinsReq)

	req, _ := http.NewRequest("POST", "http://localhost:8080/api/sendCoin", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+senderToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался статус 200, получен %d", resp.StatusCode)
	}
}

func TestBuyNotFoundItem(t *testing.T) {
	token := authenticateUser("buyer", "password")
	if token == "" {
		t.Fatal("Ошибка аутентификации покупателя")
	}

	req, _ := http.NewRequest("GET", "http://localhost:8080/api/buy/itemnotfound", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Ожидался статус 400, получен %d", resp.StatusCode)
	}
}

func authenticateUser(username, password string) string {
	authReq := map[string]string{
		"username": username,
		"password": password,
	}
	body, _ := json.Marshal(authReq)

	req, _ := http.NewRequest("POST", "http://localhost:8080/api/auth", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ""
	}

	var authResp struct {
		Token string `json:"token"`
	}
	json.NewDecoder(resp.Body).Decode(&authResp)

	return authResp.Token
}
