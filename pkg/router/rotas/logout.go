package rotas

import (
	"devbook-app/pkg/controllers"
	"net/http"
)

var rotaLogout = Rota{
	URI:            "/logout",
	Method:         http.MethodGet,
	Function:       controllers.FazerLogout,
	Authentication: true,
}
