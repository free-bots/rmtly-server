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
		MethodHandler(writer, request,
			func(writer http.ResponseWriter, request *http.Request) {

				applications := services.GetApplications()

				bytes, err := json.Marshal(applications)
				if err != nil {
					writer.WriteHeader(http.StatusBadRequest)
					return
				}

				_, _ = writer.Write(bytes)
				writer.WriteHeader(http.StatusOK)

			}, func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusForbidden)
			}, func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusForbidden)
			}, func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusForbidden)
			})
	})

	subRouter.HandleFunc("/{applicationId}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		application := services.GetApplicationById(vars["applicationId"])
		bytes, err := json.Marshal(application)
		if err != nil || application == nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		_, _ = writer.Write(bytes)
		writer.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	subRouter.HandleFunc("/run/{applicationId}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		application := services.GetApplicationById(vars["applicationId"])
		bytes, err := json.Marshal(application)
		if err != nil || application == nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		_, _ = writer.Write(bytes)
		writer.WriteHeader(http.StatusOK)

		c := make(chan bool)
		go services.RunCommand(application.Exec, c)
	})
}
