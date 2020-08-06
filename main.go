package main

import (
	"fmt"
	"log"
	"net/http"
	"rmtly-server/application/applicationUtils"
	configService "rmtly-server/config/services"
	qrService "rmtly-server/qrcode/services"
	"rmtly-server/routers"
	"rmtly-server/security/routers/routerUtils"
	"rmtly-server/security/services"
	"time"
)

//                  _   _
//                 | | | |
//   _ __ _ __ ___ | |_| |_   _ ______ ___  ___ _ ____   _____ _ __
//  | '__| '_ ` _ \| __| | | | |______/ __|/ _ \ '__\ \ / / _ \ '__|
//  | |  | | | | | | |_| | |_| |      \__ \  __/ |   \ V /  __/ |
//  |_|  |_| |_| |_|\__|_|\__, |      |___/\___|_|    \_/ \___|_|
//                         __/ |
//                        |___/

func main() {
	showBranding()
	qrService.ShowQr("rmtly-server")

	startInit()

	token, err := services.CreateJwtToken(time.Now().String())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(services.VerifyJwtToken(token))

	fmt.Println(configService.GetConfig().ImageQuality)

	startServer()
}

func startInit() {
	applicationUtils.InitIconUtils()
	configService.InitConfig()
}

func startServer() {
	router := routers.RootRouter()

	router.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		err := routerUtils.AuthenticationMiddleWare(writer, request)
		if err != nil {
			return
		}

		_, _ = writer.Write([]byte("authenticated"))
		writer.WriteHeader(http.StatusOK)
	})

	server := &http.Server{
		Addr:              "0.0.0.0:3000",
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: time.Second * 15,
		WriteTimeout:      time.Second * 15,
		IdleTimeout:       time.Second * 60,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func showBranding() {
	fmt.Println("rmtly-server running...")
	fmt.Println("//                  _   _\n" +
		"//                 | | | |\n" +
		"//   _ __ _ __ ___ | |_| |_   _ ______ ___  ___ _ ____   _____ _ __\n" +
		"//  | '__| '_ ` _ \\| __| | | | |______/ __|/ _ \\ '__\\ \\ / / _ \\ '__|\n" +
		"//  | |  | | | | | | |_| | |_| |      \\__ \\  __/ |   \\ V /  __/ |\n" +
		"//  |_|  |_| |_| |_|\\__|_|\\__, |      |___/\\___|_|    \\_/ \\___|_|\n" +
		"//                         __/ |\n" +
		"//                        |___/")
}
