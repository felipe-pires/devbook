package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa as rotas da api
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar configura as routas no router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
