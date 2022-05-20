package main

import (
	"net/http"

	"github.com/NhanNT-VNG/go-auth/internal/config"
	"github.com/NhanNT-VNG/go-auth/internal/handlers"
	"github.com/gorilla/mux"
)

func routes(config *config.Config) http.Handler {
	r := mux.NewRouter()
	passport := r.PathPrefix("/api/passport").Subrouter()

	passport.HandleFunc("/users", handlers.Repo.GetAllUser).Methods(http.MethodGet)

	return r
}
