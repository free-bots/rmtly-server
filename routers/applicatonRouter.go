package routers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rmtly-server/services"
)

const PREFIX = "/applications"

func ApplicationRouter(router *mux.Router) {
	subRouter := router.PathPrefix(PREFIX).Subrouter()

	subRouter.HandleFunc("", func(writer http.ResponseWriter, request *http.Request) {
		applications := services.GetApplications()

		bytes, err := json.Marshal(applications)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		writer.Write(bytes)
		writer.WriteHeader(http.StatusOK)
	})

}
