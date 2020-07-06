package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

const PREFIX = "/applications"

func ApplicationRouter(router *mux.Router) {
	subRouter := router.PathPrefix(PREFIX).Subrouter()

	subRouter.HandleFunc("", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("applications"))
		writer.WriteHeader(http.StatusOK)
	})

}
