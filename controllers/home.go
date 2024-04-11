package controllers

import (
	"net/http"

	"github.com/ahmed-deftoner/future-trading-bot/response"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "Future Trading Bot")
}
