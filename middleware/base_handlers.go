package middleware

import (
	"encoding/json"
	"net/http"
	"simulator/env"
	"simulator/logger"
	"simulator/middleware/api-messages"
	"time"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	// TODO: handle Encode errors
	json.NewEncoder(w).Encode(api_messages.BasicResponse{
		Time:    time.Now(),
		Message: "pong",
	})
	logger.Info.Println("ping OK, port "+env.GetPortStr(), r.RemoteAddr)
}
