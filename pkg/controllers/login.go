package controllers

import (
	"devbook-app/pkg/utils"
	"net/http"
)

// CarregarTelaDeLogin() vai renderizar a tela de login.
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}
