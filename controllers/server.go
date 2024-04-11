package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ahmed-deftoner/future-trading-bot/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		log.Printf("Cannot connect to the %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		log.Printf("Connected to the %s database", Dbdriver)
	}

	server.DB.AutoMigrate(&models.Key{})
}

func (server *Server) Run(addr string) {
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe("127.0.0.1"+addr, server.Router))
}
