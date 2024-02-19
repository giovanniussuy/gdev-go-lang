package server

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/giovanniussuy/gdev-go-lang/infra/config"
	"github.com/giovanniussuy/gdev-go-lang/infra/router"
)

type (
	Server struct {
		appConfigurations *config.AppConfigurations
		serverEngine      router.Server
	}
)

func CreateServerStruct() *Server {
	return &Server{}
}

func (serverStruct *Server) LoadEnvConfiguration() *Server {
	appConfig, err := config.LoadAppConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error to load app configuration!"), err)
	}

	serverStruct.appConfigurations = appConfig
	return serverStruct
}

func (serverStruct *Server) StartWebServerEngine() *Server {
	intParsed, err := strconv.ParseInt(serverStruct.appConfigurations.ServerPort, 10, 64)
	if err != nil {
		log.Fatal("Error to parse the Port type!", err)
	}

	durationParser, err := time.ParseDuration(serverStruct.appConfigurations.ServerTimeout + "s")
	if err != nil {
		log.Fatal("Error to parse the Timeout duration!", err)
	}

	server := router.
		StartWebEngine().
		WithAppName().
		WithPort(intParsed).
		WithDuration(durationParser)

	fmt.Println("INFO: Router server has been successfully configured.")

	serverStruct.serverEngine = server
	return serverStruct
}

func (serverStruct *Server) Listen() {
	serverStruct.serverEngine.Listen()
}
