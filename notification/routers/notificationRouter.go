package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"rmtly-server/contants"
	"rmtly-server/notification/services"
	"rmtly-server/security/routers/routerUtils"
)

const PREFIX = "/notifications"

func NotificationRouter(router *mux.Router) {
	subRouter := router.PathPrefix(PREFIX).Subrouter()

	if !contants.IS_DEV_MODE {
		subRouter.Use(routerUtils.AuthorizationMiddleware)
	}

	subRouter.HandleFunc("", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		services.SendAsync(request.FormValue("title"), request.FormValue("message"))
	})
}
