package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"rmtly-server/application/routers"
	routers2 "rmtly-server/notification/routers"
	routers3 "rmtly-server/security/routers"
)

const BRANDING = "                _   _\n" +
	"               | | | |\n" +
	" _ __ _ __ ___ | |_| |_   _ ______ ___  ___ _ ____   _____ _ __\n" +
	"| '__| '_ ` _ \\| __| | | | |______/ __|/ _ \\ '__\\ \\ / / _ \\ '__|\n" +
	"| |  | | | | | | |_| | |_| |      \\__ \\  __/ |   \\ V /  __/ |\n" +
	"|_|  |_| |_| |_|\\__|_|\\__, |      |___/\\___|_|    \\_/ \\___|_|\n" +
	"                       __/ |\n" +
	"                      |___/                                     "

func RootRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(BRANDING))
		writer.WriteHeader(http.StatusOK)
	})

	routers.ApplicationRouter(router)
	routers2.NotificationRouter(router)
	routers3.AuthenticationRouter(router)

	return router
}
