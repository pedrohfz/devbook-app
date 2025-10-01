package controllers

import (
	"devbook-app/internal/config"
	"devbook-app/internal/cookies"
	"devbook-app/internal/request"
	"devbook-app/pkg/models"
	"devbook-app/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CarregarTelaDeLogin() vai renderizar a tela de login.
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

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

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutarTemplate(w, "home.html", struct {
		Publicacoes []models.Publicacao
		UsuarioID   uint64
	}{
		Publicacoes: publicacoes,
		UsuarioID:   usuarioID,
	})
}

// CarregarPaginaDeAtualizacaoDePublicacao() carrega a página de edição de publicação.
func CarregarPaginaDeAtualizacaoDePublicacao(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(param["publicacaoID"], 10, 64)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.ErroAPI{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d", config.APIURL, publicacaoID)
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

	var publicacao models.Publicacao
	if err = json.NewDecoder(response.Body).Decode(&publicacao); err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, utils.ErroAPI{Erro: err.Error()})
		return
	}

	utils.ExecutarTemplate(w, "atualizar-publicacao.html", publicacao)
}

// CarregarPaginaDeUsuarios() carrega a página com os usuários que atendem o filtro passado.
func CarregarPaginaDeUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.APIURL, nomeOuNick)

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

	var usuarios []models.Usuario
	if err = json.NewDecoder(response.Body).Decode(&usuarios); err != nil {
		utils.JSON(w, http.StatusUnprocessableEntity, utils.ErroAPI{Erro: err.Error()})
		return
	}

	utils.ExecutarTemplate(w, "usuarios.html", usuarios)
}

// CarregarPerfilDoUsuario() carrega a página do perfil do usuário.
func CarregarPerfilDoUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(param["usuarioID"], 10, 64)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.ErroAPI{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioLogadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if usuarioID == usuarioLogadoID {
		http.Redirect(w, r, "/perfil", 302)
		return
	}

	usuario, err := models.BuscarUsuarioCompleto(usuarioID, r)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, utils.ErroAPI{Erro: err.Error()})
		return
	}

	utils.ExecutarTemplate(w, "usuario.html", struct {
		Usuario         models.Usuario
		UsuarioLogadoID uint64
	}{
		Usuario:         usuario,
		UsuarioLogadoID: usuarioLogadoID,
	})
}

// CarregarPerfilDoUsuarioLogado() carrega a página do perfil do usuário logado.
func CarregarPerfilDoUsuarioLogado(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	usuario, err := models.BuscarUsuarioCompleto(usuarioID, r)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, utils.ErroAPI{Erro: err.Error()})
		return
	}

	utils.ExecutarTemplate(w, "perfil.html", usuario)
}

// CarregarPaginaDeEdicaoDeUsuario() carrega a página para edição dos dados do usuário.
func CarregarPaginaDeEdicaoDeUsuario(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	canal := make(chan models.Usuario)
	go models.BuscarDadosDoUsuario(canal, usuarioID, r)
	usuario := <-canal

	if usuario.ID == 0 {
		utils.JSON(w, http.StatusInternalServerError, utils.ErroAPI{Erro: "erro ao buscar o usuário"})
		return
	}

	utils.ExecutarTemplate(w, "editar-usuario.html", usuario)
}

// CarregarPaginaDeAtualizacaoDeSenha() carrega a página para atualização da senha do usuário.
func CarregarPaginaDeAtualizacaoDeSenha(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "atualizar-senha.html", nil)
}
