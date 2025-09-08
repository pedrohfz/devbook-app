package controllers

import (
	"bytes"
	"devbook-app/internal/config"
	"devbook-app/internal/request"
	"devbook-app/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// CriarPublicacao() chama a API para cadastrar uma publicação no banco de dados.
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publicacao, err := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, err := request.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, utils.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		utils.TratarStatusCodeDeErro(w, response)
		return
	}

	utils.JSON(w, response.StatusCode, nil)
}
