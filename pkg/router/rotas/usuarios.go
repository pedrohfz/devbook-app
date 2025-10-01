package rotas

import (
	"devbook-app/pkg/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:            "/criar-usuario",
		Method:         http.MethodGet,
		Function:       controllers.CarregarPaginaDeCadastroDeUsuario,
		Authentication: false,
	},
	{
		URI:            "/usuarios",
		Method:         http.MethodPost,
		Function:       controllers.CriarUsuario,
		Authentication: false,
	},
	{
		URI:            "/buscar-usuarios",
		Method:         http.MethodGet,
		Function:       controllers.CarregarPaginaDeUsuarios,
		Authentication: true,
	},
	{
		URI:            "/usuarios/{usuarioID}",
		Method:         http.MethodGet,
		Function:       controllers.CarregarPerfilDoUsuario,
		Authentication: true,
	},
	{
		URI:            "/usuarios/{usuarioID}/parar-de-seguir",
		Method:         http.MethodPost,
		Function:       controllers.PararDeSeguirUsuario,
		Authentication: true,
	},
	{
		URI:            "/usuarios/{usuarioID}/seguir",
		Method:         http.MethodPost,
		Function:       controllers.SeguirUsuario,
		Authentication: true,
	},
	{
		URI:            "/perfil",
		Method:         http.MethodGet,
		Function:       controllers.CarregarPerfilDoUsuarioLogado,
		Authentication: true,
	},
	{
		URI:            "/editar-usuario",
		Method:         http.MethodGet,
		Function:       controllers.CarregarPaginaDeEdicaoDeUsuario,
		Authentication: true,
	},
	{
		URI:            "/editar-usuario",
		Method:         http.MethodPut,
		Function:       controllers.EditarUsuario,
		Authentication: true,
	},
	{
		URI:            "/atualizar-senha",
		Method:         http.MethodGet,
		Function:       controllers.CarregarPaginaDeAtualizacaoDeSenha,
		Authentication: true,
	},
	{
		URI:            "/atualizar-senha",
		Method:         http.MethodPost,
		Function:       controllers.AtualizarSenha,
		Authentication: true,
	},
}
