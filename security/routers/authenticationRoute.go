package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	qrService "rmtly-server/qrcode/services"
	"rmtly-server/routers/routersUtil"
	"rmtly-server/security/interfaces"
	"rmtly-server/security/services"
)

const PREFIX = "/authentication"

func AuthenticationRouter(router *mux.Router) {
	subRouter := router.PathPrefix(PREFIX).Subrouter()

	subRouter.HandleFunc("/signUp", func(writer http.ResponseWriter, request *http.Request) {
		routersUtil.ContentTypeJson(writer)

		defer request.Body.Close()

		signUpRequest := interfaces.SignUpRequest{}
		err := json.NewDecoder(request.Body).Decode(signUpRequest)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
		}

		token, err := services.CreateJwtToken(signUpRequest)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
		}

		_, _ = writer.Write([]byte(token))
		writer.WriteHeader(http.StatusOK)

	}).Methods(http.MethodPost)

	subRouter.HandleFunc("/code", func(writer http.ResponseWriter, request *http.Request) {
		code := "authenticationCode"
		qrService.ShowQr(code)
		fmt.Printf("Scan the code or type: %s", code)
		writer.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

}
