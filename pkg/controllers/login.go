package controllers

import (
	"bytes"
	"devbook-app/internal/config"
	"devbook-app/pkg/models"
	"devbook-app/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// FazerLogin() utiliza o e-mail e senha do usuário para autenticar na aplicação.
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, utils.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		utils.TratarStatusCodeDeErro(w, response)
		return
	}

	var dadosAuth models.DadosAuth
	if err = json.NewDecoder(response.Body).Decode(&dadosAuth); err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, utils.ErroAPI{Erro: err.Error()})
		return
	}

	// TODO: Cookies.

	utils.JSON(w, http.StatusOK, nil)
}
