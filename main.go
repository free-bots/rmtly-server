package main

import (
	"fmt"
	"log"
	"net/http"
	"rmtly-server/applicationUtils"
	"rmtly-server/routers"
	"rmtly-server/services"
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
	fmt.Println("rmtly-server running...")

	router := routers.RootRouter()

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

	const path = "./test.desktop"
	applicationEntry := applicationUtils.Parse(path, true)

	c := make(chan bool)
	go services.RunCommand(applicationEntry.Exec, c)

	fmt.Printf("running %s succesful %t", applicationEntry.Exec, <-c)
}
