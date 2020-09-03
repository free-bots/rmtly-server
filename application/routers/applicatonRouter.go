package routers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rmtly-server/application/services"
	notificationService "rmtly-server/notification/services"
	"rmtly-server/routers/routersUtil"
	"rmtly-server/security/routers/routerUtils"
	"strconv"
)

const PREFIX = "/applications"

func ApplicationRouter(router *mux.Router) {
	subRouter := router.PathPrefix(PREFIX).Subrouter()

	subRouter.Use(routerUtils.AuthorizationMiddleware)

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

				icon, _ := strconv.ParseBool(request.FormValue("icon"))

				applications := services.GetApplications(icon)

				// todo if query icon -> merge icon

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
		icon, _ := strconv.ParseBool(request.FormValue("icon"))
		application := services.GetApplicationById(vars["applicationId"], icon)
		// if query icon -> merge icon
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
		vars := mux.Vars(request)
		icon := services.GetIconOfApplication(vars["applicationId"])
		bytes, err := json.Marshal(icon)
		if err != nil || icon == nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		routersUtil.ContentTypeJson(writer)

		_, _ = writer.Write(bytes)
		writer.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	subRouter.HandleFunc("/run/{applicationId}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		icon, _ := strconv.ParseBool(request.FormValue("icon"))
		application := services.GetApplicationById(vars["applicationId"], icon)
		bytes, err := json.Marshal(application)
		if err != nil || application == nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		routersUtil.ContentTypeJson(writer)

		_, _ = writer.Write(bytes)

		writer.WriteHeader(http.StatusOK)

		c := make(chan bool)
		go services.RunCommand(application.Exec, c)
		notificationService.SendAsync(application.Name, "executed by rmtly-server")
	})
}
