package routerUtils

import (
	"fmt"
	"net/http"
	"rmtly-server/security/services"
	"strings"
)

func AuthenticationMiddleWare(writer http.ResponseWriter, request *http.Request) error {
	if authenticationHelper(writer, request) {
		return nil
	}
	return fmt.Errorf("unauthorized")

}

func authenticationHelper(writer http.ResponseWriter, request *http.Request) (authorized bool) {
	bearerToken := request.Header.Get("Authorization")

	tokenArray := strings.Split(bearerToken, " ")

	if len(tokenArray) <= 1 || tokenArray[0] != "Bearer" || tokenArray[1] == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		return false
	}

	authorizationToken := tokenArray[1]

	deviceId, err := services.VerifyJwtToken(authorizationToken)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusUnauthorized)
		return false
	}

	fmt.Printf("request from device: %s", deviceId)
	return true
}
