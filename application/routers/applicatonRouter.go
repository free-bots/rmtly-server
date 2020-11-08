package routers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rmtly-server/application/services"
	"rmtly-server/routers/routersUtil"
	"strconv"
)

const PREFIX = "/applications"

func ApplicationRouter(router *mux.Router) {
	subRouter := router.PathPrefix(PREFIX).Subrouter()

	//subRouter.Use(routerUtils.AuthorizationMiddleware)

	subRouter.Queries("sortedBy", "{*}").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		routersUtil.ContentTypeJson(writer)
		sortedBy := request.FormValue("sortedBy")
		if sortedBy == "" {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		sortedResponse := services.GetApplicationsSortedBy(sortedBy)
		bytes, err := json.Marshal(sortedResponse)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		_, _ = writer.Write(bytes)

		writer.WriteHeader(http.StatusOK)

	}).Methods(http.MethodGet)

	subRouter.HandleFunc("", func(writer http.ResponseWriter, request *http.Request) {
		routersUtil.MethodHandler(writer, request,
			func(writer http.ResponseWriter, request *http.Request) {
				routersUtil.ContentTypeJson(writer)

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
		applicationId := getApplicationId(request)
		application := services.GetApplicationById(applicationId)

		bytes, err := json.Marshal(application)
		if err != nil || application == nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		routersUtil.ContentTypeJson(writer)

		_, _ = writer.Write(bytes)
		writer.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	subRouter.HandleFunc("/{applicationId}/icon", func(writer http.ResponseWriter, request *http.Request) {
		applicationId := getApplicationId(request)

		iconBuffer := services.GetApplicationIcon(applicationId)
		if iconBuffer == nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		writer.Header().Set("Content-Type", "image/png")
		writer.Header().Set("Content-Length", strconv.Itoa(len(iconBuffer.Bytes())))
		_, _ = writer.Write(iconBuffer.Bytes())
		writer.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	subRouter.HandleFunc("/{applicationId}/execute", func(writer http.ResponseWriter, request *http.Request) {
		applicationId := getApplicationId(request)

		response := services.Execute(applicationId)

		bytes, err := json.Marshal(response)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		routersUtil.ContentTypeJson(writer)

		_, _ = writer.Write(bytes)

		writer.WriteHeader(http.StatusOK)
	})
}

func getApplicationId(request *http.Request) string {
	vars := mux.Vars(request)
	return vars["applicationId"]
}
