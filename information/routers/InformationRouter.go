package routers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rmtly-server/information/services"
)

const PREFIX = "/information"

func InformationRouter(router *mux.Router) {
	subRouter := router.PathPrefix(PREFIX).Subrouter()

	subRouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		bytes, err := json.Marshal(services.GetInformation())
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		_, _ = writer.Write(bytes)
		writer.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
}
