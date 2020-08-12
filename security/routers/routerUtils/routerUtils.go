package routerUtils

import (
	"fmt"
	"net/http"
	"rmtly-server/security/services"
	"strings"
)

func AuthorizationMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		bearerToken := request.Header.Get("Authorization")

		tokenArray := strings.Split(bearerToken, " ")

		if len(tokenArray) <= 1 || tokenArray[0] != "Bearer" || tokenArray[1] == "" {
			writer.WriteHeader(http.StatusUnauthorized)
			http.Error(writer, "Authorization header missing or invalid", http.StatusUnauthorized)
			return
		}

		authorizationToken := tokenArray[1]

		deviceId, err := services.VerifyJwtToken(authorizationToken)
		if err != nil {
			fmt.Println(err)
			writer.WriteHeader(http.StatusUnauthorized)
			http.Error(writer, "invalid token", http.StatusUnauthorized)
			return
		}

		fmt.Printf("request from device: %s", deviceId)
		handler.ServeHTTP(writer, request)
	})
}
