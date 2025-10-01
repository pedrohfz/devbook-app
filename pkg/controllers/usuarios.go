package controllers

import (
	"bytes"
	"devbook-app/internal/config"
	"devbook-app/internal/cookies"
	"devbook-app/internal/request"
	"devbook-app/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CriarUsuario chama a API para cadastrar um usuário no banco de dados.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios", config.APIURL)
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

	utils.JSON(w, response.StatusCode, nil)
}

// PararDeSeguirUsuario() chama a API para parar de seguir um usuário.
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(param["usuarioID"], 10, 64)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/parar-de-seguir", config.APIURL, usuarioID)
	response, err := request.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
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

// SeguirUsuario() chama a API para seguir um usuário.
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(param["usuarioID"], 10, 64)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/seguir", config.APIURL, usuarioID)
	response, err := request.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
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

// EditarUsuario() chama a API para editar um usuário.
func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
	})
	
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.ErroAPI{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)
	response, err := request.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(usuario))
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
