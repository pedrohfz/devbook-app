package controllers

import (
	"devbook-app/internal/config"
	"devbook-app/internal/request"
	"devbook-app/pkg/utils"
	"fmt"
	"net/http"
)

// CarregarTelaDeLogin() vai renderizar a tela de login.
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaDeCadastroDeUsuario() vai carregar a página de cadastro de usuário.
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// CarregarPaginaPrincipal() vai carregar a página principal com as publicações.
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, err := request.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		// TODO: Next class.
	}

	fmt.Println(response)
	utils.ExecutarTemplate(w, "home.html", nil)
}
