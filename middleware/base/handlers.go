package base

import (
	"encoding/json"
	"net/http"
	"simulator/env"
	"simulator/logger"
	"time"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(BasicResponse{
		Time:    time.Now(),
		Message: "pong",
	})
	if ErrorHandling(err, "ping error", w) != nil {
		return
	}
	logger.Info.Println("ping OK, port "+env.GetPortStr(), r.RemoteAddr)
}

func ErrorHandling(err error, msg string, w http.ResponseWriter) error {
	if err != nil {
		logger.Error.Println(err, ":", msg)
		errRes := ErrorResponse{
			ErrorType: err.Error(),
			Message:   "error: " + msg,
		}
		fatalErr := json.NewEncoder(w).Encode(errRes)
		if fatalErr != nil {
			panic(fatalErr)
		}
	}
	return err
}
