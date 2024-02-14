package main

import "github.com/giovanniussuy/gdev-go-lang/infra/server"

func main() {
	server.
		CreateServerStruct().
		LoadEnvConfiguration().
		StartWebServerEngine().
		Listen()
}
