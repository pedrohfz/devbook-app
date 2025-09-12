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
	{
		URI:            "/publicacoes/{publicacaoID}/curtir",
		Method:         http.MethodPost,
		Function:       controllers.CurtirPublicacao,
		Authentication: true,
	},
}
