package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"rmtly-server/application/applicationUtils"
	"rmtly-server/config/interfaces"
	configService "rmtly-server/config/services"
	"rmtly-server/routers"
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
	startInit()
	config := configService.GetConfig()
	startServer(config)
}

func startInit() {
	applicationUtils.InitIconUtils()
	configService.InitConfig()
}

func startServer(config interfaces.Config) {
	router := routers.RootRouter()

	handler := handlers.CORS(
		handlers.AllowedHeaders([]string{
			"Accept",
			"Accept-Language",
			"Content-Type",
			"Content-Language",
			"Origin",
		}),
	)(router)

	server := &http.Server{
		Addr:              config.Network.Address,
		Handler:           handler,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: time.Second * 150,
		WriteTimeout:      time.Second * 150,
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
	fmt.Println("//                  _   _\n"+
		"//                 | | | |\n"+
		"//   _ __ _ __ ___ | |_| |_   _ ______ ___  ___ _ ____   _____ _ __\n"+
		"//  | '__| '_ ` _ \\| __| | | | |______/ __|/ _ \\ '__\\ \\ / / _ \\ '__|\n"+
		"//  | |  | | | | | | |_| | |_| |      \\__ \\  __/ |   \\ V /  __/ |\n"+
		"//  |_|  |_| |_| |_|\\__|_|\\__, |      |___/\\___|_|    \\_/ \\___|_|\n"+
		"//                         __/ |\n"+
		"//                        |___/\n",
		"rmtly-server running...")
}
