package base

import (
	"encoding/json"
	"net/http"
	"simulator/env"
	"simulator/logger"
	"time"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	// TODO: handle Encode errors
	json.NewEncoder(w).Encode(BasicResponse{
		Time:    time.Now(),
		Message: "pong",
	})
	logger.Info.Println("ping OK, port "+env.GetPortStr(), r.RemoteAddr)
}
