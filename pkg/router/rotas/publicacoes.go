package rotas

import (
	"devbook-app/pkg/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:            "/publicacoes",
		Method:         http.MethodPost,
		Function:       controllers.CriarPublicacao,
		Authentication: true,
	},
}
