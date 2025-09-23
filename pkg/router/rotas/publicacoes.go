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
	{
		URI:            "/publicacoes/{publicacaoID}/descurtir",
		Method:         http.MethodPost,
		Function:       controllers.DescurtirPublicacao,
		Authentication: true,
	},
	{
		URI:            "/publicacoes/{publicacaoID}/atualizar",
		Method:         http.MethodGet,
		Function:       controllers.CarregarPaginaDeAtualizacaoDePublicacao,
		Authentication: true,
	},
	{
		URI:            "/publicacoes/{publicacaoID}",
		Method:         http.MethodPut,
		Function:       controllers.AtualizarPublicacao,
		Authentication: true,
	},
	{
		URI:            "/publicacoes/{publicacaoID}",
		Method:         http.MethodDelete,
		Function:       controllers.DeletarPublicacao,
		Authentication: true,
	},
}
