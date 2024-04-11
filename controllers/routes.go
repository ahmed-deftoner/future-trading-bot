package controllers

import "github.com/ahmed-deftoner/future-trading-bot/middleware"

func (r *Server) initializeRoutes() {
	s := r.Router.PathPrefix("/bot").Subrouter()
	s.HandleFunc("/", middleware.MiddlewareJSON(r.Home)).Methods("GET")
	s.HandleFunc("/create-key", middleware.MiddlewareJSON(r.CreateKey)).Methods("POST")
}
