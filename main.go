package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rmtly-server/application/applicationUtils"
	"rmtly-server/application/applicationUtils/parser/application"
	configService "rmtly-server/config/services"
	qrService "rmtly-server/qrcode/services"
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
	qrService.ShowQr("rmtly-server")

	startInit()

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

		const path = "./test.desktop"
		applicationEntry := application.Parse(path, true)

		//c := make(chan bool)
		//go services.RunCommand(applicationEntry.Exec, c)
		//
		//fmt.Printf("running %s succesful %t", applicationEntry.Exec, <-c)
		bytes, _ := json.Marshal(applicationEntry)
		_, _ = writer.Write(bytes)
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
