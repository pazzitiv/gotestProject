package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"gotestProject/app/common"
	"log"
	"net/http"
)

var Server *mux.Router

func RunServer() {
	Server = mux.NewRouter().StrictSlash(true)
	Server.HandleFunc("/metrics", metricsHandler)

	serverAddress := fmt.Sprintf("%s:%d", common.App.Server.Host, common.App.Server.Port)
	fmt.Printf("Server is listening on %s...\n", serverAddress)
	log.Println("API started.")
	err := http.ListenAndServe(serverAddress, Server)
	if err != nil {
		log.Fatalf("[FATAL] %s", err.Error())
	}
}
