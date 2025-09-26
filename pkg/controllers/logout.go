package controllers

import (
	"devbook-app/internal/cookies"
	"net/http"
)

// FazerLogout() remove os dados de autenticação no browser do usuário.
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", 302)
}
