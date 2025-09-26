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
}
