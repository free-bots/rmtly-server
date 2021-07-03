package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"rmtly-server/application/interfaces"
	"rmtly-server/contants"
	services2 "rmtly-server/scripts/services"
	"rmtly-server/security/routers/routerUtils"
)

const PREFIX = "/scripts"

func ScriptRouter(router *mux.Router) {
	subRouter := router.PathPrefix(PREFIX).Subrouter()

	if !contants.IS_DEV_MODE {
		subRouter.Use(routerUtils.AuthorizationMiddleware)
	}

	subRouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		bytes, err := json.Marshal(services2.GetAllScripts())
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		_, _ = writer.Write(bytes)
		writer.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	subRouter.HandleFunc("/{scriptName}/execute", func(writer http.ResponseWriter, request *http.Request) {
		scriptName := getScriptName(request)

		defer func() {
			_ = request.Body.Close()
		}()

		executeRequest := new(interfaces.ExecuteRequest)
		err := json.NewDecoder(request.Body).Decode(executeRequest)

		if err != nil {
			fmt.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		services2.Exec(scriptName, *executeRequest)

		writer.WriteHeader(http.StatusOK)
	}).Methods(http.MethodPost)
}

func getScriptName(request *http.Request) string {
	vars := mux.Vars(request)
	return vars["scriptName"]
}
