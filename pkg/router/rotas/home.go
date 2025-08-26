package rotas

import (
	"devbook-app/pkg/controllers"
	"net/http"
)

var rotaPaginaPrincipal = Rota{
	URI:            "/home",
	Method:         http.MethodGet,
	Function:       controllers.CarregarPaginaPrincipal,
	Authentication: true,
}
