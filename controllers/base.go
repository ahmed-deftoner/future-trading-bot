package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/ahmed-deftoner/future-trading-bot/bitget"
	"github.com/ahmed-deftoner/future-trading-bot/codec"
	"github.com/ahmed-deftoner/future-trading-bot/models"
	"github.com/joho/godotenv"
)

var server = Server{}

func Run() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error getting env")
	} else {
		log.Println("Getting Values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	placeTrade()
	server.Run(":8080")
}

func decryptKeyInfo(api_key string, secret_key string, passphrase string) (string, string, string) {
	decrypted_api_key, err := codec.Decrypt(api_key)
	if err != nil {
		log.Fatal("Error Decrypting API key")
	}

	decrypted_secret_key, err := codec.Decrypt(secret_key)
	if err != nil {
		log.Fatal("Error Decrypting API key")
	}

	decrypted_passphrase, err := codec.Decrypt(passphrase)
	if err != nil {
		log.Fatal("Error Decrypting API key")
	}

	return decrypted_api_key, decrypted_secret_key, decrypted_passphrase
}

func placeTrade() {
	keyModel := &models.Key{}
	key, err := keyModel.FindKeyByEmail(server.DB, "ahmedghtwhts786@gmail.com")
	if err != nil {
		log.Fatal("Error Getting key")
	}

	fmt.Println(key.UserEmail)

	api_key, secret_key, passphrase := decryptKeyInfo(key.ApiKey, key.SecretKey, key.Passphrase)
	longOrder := bitget.OrderRequest{
		Size:      "10",
		Side:      "close_long",
		OrderType: "Market",
	}

	shortOrder := bitget.OrderRequest{
		Size:      "10",
		Side:      "close_short",
		OrderType: "Market",
	}

	batchOrderRequest := bitget.BitgetBatchOrderRequest{
		Symbol:        "SXRPSUSDT_SUMCBL",
		MarginCoin:    "SUSDT",
		OrderDataList: []bitget.OrderRequest{longOrder, shortOrder},
	}

	resp, err := bitget.PlaceBitgetBatchOrder(api_key, secret_key, passphrase, &batchOrderRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
