package routersUtil

import "net/http"

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

func ContentTypeJson(writer http.ResponseWriter) {
	writer.Header().Add("content-type", "application/json")
}
