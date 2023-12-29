package main

import (
	"PartyRoom.API/internal/config"
	"PartyRoom.API/internal/http-server/handlers/tag/createTag"
	"PartyRoom.API/internal/http-server/handlers/tag/deleteTag"
	"PartyRoom.API/internal/http-server/handlers/tag/updateTag"
	"PartyRoom.API/internal/http-server/handlers/user/profile"
	"PartyRoom.API/internal/http-server/handlers/user/refreshToken"
	registrationUser "PartyRoom.API/internal/http-server/handlers/user/registration"
	"PartyRoom.API/internal/http-server/handlers/user/signIn"
	middleware "PartyRoom.API/internal/http-server/middleware/desirializeUser"
	"PartyRoom.API/internal/service/authService"
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
	authServ := authService.New(storage)
	router := chi.NewRouter()

	router.Post("/register", registrationUser.New(storage))
	router.Route("/auth", func(r chi.Router) {
		r.Post("/refresh", refreshToken.New(&authServ, cfg))
		r.Post("/signIn", signIn.New(&authServ))
	})

	router.Route("/profile", func(r chi.Router) {
		r.Use(middleware.ValidateJwt)
		r.Get("/", profile.New(storage))
	})

	router.Route("/tag", func(r chi.Router) {
		r.Use(middleware.ValidateJwt)
		r.Post("/", createTag.New(storage))
		r.Put("/", updateTag.New(storage))
		r.Delete("/{tagID}", deleteTag.New(storage))
	})
	err = http.ListenAndServe("localhost:8080", router)
	if err != nil {
		return
	}
}
