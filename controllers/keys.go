package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ahmed-deftoner/future-trading-bot/codec"
	"github.com/ahmed-deftoner/future-trading-bot/models"
	"github.com/ahmed-deftoner/future-trading-bot/response"
)

type KeyRequest struct {
	ApiKey          string  `json:"api_key"`
	SecretKey       string  `json:"secret_key"`
	Passphrase      string  `json:"passphrase"`
	UserEmail       string  `json:"user_email"`
	TradeAmount     int     `json:"trade_amount"`
	AllowedCoins    int     `json:"allowed_coins"`
	CapitalPerTrade float64 `json:"capital_per_trade"`
	Start           bool    `json:"start"`
}

func (server *Server) CreateKey(w http.ResponseWriter, r *http.Request) {
	var keyRequest KeyRequest
	err := json.NewDecoder(r.Body).Decode(&keyRequest)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("bad params for key"))
		return
	}

	encryptedApiKey, err := codec.Encrypt(keyRequest.ApiKey)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, errors.New("failed to encrypt api key"))
		return
	}
	encryptedSecretKey, err := codec.Encrypt(keyRequest.SecretKey)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, errors.New("failed to encrypt secret key"))
		return
	}
	encryptedPassphrase, err := codec.Encrypt(keyRequest.Passphrase)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, errors.New("failed to encrypt passphrase"))
		return
	}

	// Create Key object
	key := &models.Key{
		ApiKey:          encryptedApiKey,
		SecretKey:       encryptedSecretKey,
		Passphrase:      encryptedPassphrase,
		UserEmail:       keyRequest.UserEmail,
		TradeAmount:     keyRequest.TradeAmount,
		AllowedCoins:    keyRequest.AllowedCoins,
		CapitalPerTrade: keyRequest.CapitalPerTrade,
		Start:           keyRequest.Start,
	}

	key, err = key.SaveKey(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, errors.New("failed to save key"))
		return
	}

	response.JSON(w, http.StatusOK, key)
}
