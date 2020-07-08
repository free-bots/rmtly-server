package routers

import (
	"github.com/gorilla/mux"
	"net/http"
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
		writer.Write([]byte(BRANDING))
		writer.WriteHeader(http.StatusOK)
	})

	ApplicationRouter(router)

	return router
}

func MethodHandler(writer http.ResponseWriter, request *http.Request,
	onGet func(writer http.ResponseWriter, request *http.Request),
	onPost func(writer http.ResponseWriter, request *http.Request),
	onPut func(writer http.ResponseWriter, request *http.Request),
	onDelete func(writer http.ResponseWriter, request *http.Request)) {
	switch request.Method {
	case http.MethodGet:
		onGet(writer, request)
	case http.MethodPost:
		onPost(writer, request)
	case http.MethodPut:
		onPut(writer, request)
	case http.MethodDelete:
		onDelete(writer, request)
	default:
		writer.WriteHeader(http.StatusBadRequest)
	}
}
