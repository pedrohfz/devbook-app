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
}
