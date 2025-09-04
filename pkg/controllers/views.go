package controllers

import (
	"devbook-app/internal/config"
	"devbook-app/internal/request"
	"devbook-app/pkg/models"
	"devbook-app/pkg/utils"
	"encoding/json"
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
		utils.JSON(w, http.StatusInternalServerError, utils.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		utils.TratarStatusCodeDeErro(w, response)
		return
	}

	var publicacoes []models.Publicacao
	if err = json.NewDecoder(response.Body).Decode(&publicacoes); err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, utils.ErroAPI{Erro: err.Error()})
		return
	}

	utils.ExecutarTemplate(w, "home.html", publicacoes)
}
