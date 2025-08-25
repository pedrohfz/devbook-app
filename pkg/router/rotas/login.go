package rotas

import (
	"devbook-app/pkg/controllers"
	"net/http"
)

var rotasLogin = []Rota{
	{
		URI:            "/",
		Method:         http.MethodGet,
		Function:       controllers.CarregarTelaDeLogin,
		Authentication: false,
	},
	{
		URI:            "/login",
		Method:         http.MethodGet,
		Function:       controllers.CarregarTelaDeLogin,
		Authentication: false,
	},
	{
		URI:            "/login",
		Method:         http.MethodPost,
		Function:       controllers.FazerLogin,
		Authentication: false,
	},
}
