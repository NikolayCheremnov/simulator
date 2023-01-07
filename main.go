package main

import (
	"log"
	"net/http"
	"simulator/env"
	"simulator/logger"
	"simulator/router"
)

func main() {
	env.Init()
	port := env.GetPortStr()
	r := router.Router()
	logger.Info.Println("Starting server on the port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
