package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da aplicação web.
type Rota struct {
	URI            string
	Method         string
	Function       func(w http.ResponseWriter, r *http.Request)
	Authentication bool
}

// Configurar() coloca todas as rotas da aplicação dentro do router.
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotaPaginaPrincipal)

	for _, rota := range rotas {
		router.HandleFunc(rota.URI, rota.Function).Methods(rota.Method)
	}

	fileServer := http.FileServer(http.Dir("internal/assets"))
	router.PathPrefix("/internal/assets/").Handler(http.StripPrefix("/internal/assets/", fileServer))

	return router
}
