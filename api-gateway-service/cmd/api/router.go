package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
)

func CreateRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(TokenAuth))
		router.Use(jwtauth.Authenticator)

	})

	router.Group(func(router chi.Router) {
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("User Service is online"))
		})
	})

	return router
}
