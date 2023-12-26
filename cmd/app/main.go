package main

import (
	"PartyRoom.API/internal/config"
	"PartyRoom.API/internal/http-server/handlers/user/profile"
	"PartyRoom.API/internal/http-server/handlers/user/refreshToken"
	registrationUser "PartyRoom.API/internal/http-server/handlers/user/registration"
	"PartyRoom.API/internal/http-server/handlers/user/signIn"
	middleware "PartyRoom.API/internal/http-server/middleware/desirializeUser"
	"PartyRoom.API/internal/storage/postgresql"
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.New(".")
	if err != nil {
		log.Fatalln("Failed to load environment variavles! \n", err.Error())
	}
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.DBHost, cfg.DBUserName, cfg.DBUserPassword, cfg.DBName, cfg.DBPort)
	storage, err := postgresql.New(connectionString)
	if err != nil {
		log.Fatalln("failed to init storage \n", err.Error())
	}

	router := chi.NewRouter()

	router.Post("/signIn", signIn.New(storage, storage))
	router.Post("/register", registrationUser.New(storage))
	router.Route("/auth", func(r chi.Router) {
		r.Post("/refresh", refreshToken.New(storage, storage, cfg))
	})
	router.Route("/profile", func(r chi.Router) {
		r.Use(middleware.ValidateJwt)
		r.Get("/", profile.New(storage))
	})
	err = http.ListenAndServe("localhost:8080", router)
	if err != nil {
		return
	}
}
