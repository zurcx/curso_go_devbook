package rotas

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas a aplicacao web
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotaPaginaPrincipal)
	rotas = append(rotas, rotasPublicacoes...)

	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			router.HandleFunc(rota.URI,
				middlewares.Autenticar(middlewares.Logger(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.URI,
				middlewares.Logger(rota.Funcao),
			).Methods(rota.Metodo)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	return router
}
