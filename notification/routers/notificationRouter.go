package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"rmtly-server/notification/services"
)

const PREFIX = "/notifications"

func NotificationRouter(router *mux.Router) {
	subRouter := router.PathPrefix(PREFIX).Subrouter()

	subRouter.HandleFunc("", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		services.SendAsync(request.FormValue("title"), request.FormValue("message"))
	})
}
