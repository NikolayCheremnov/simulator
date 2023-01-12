package base

import (
	"time"
)

type BasicResponse struct {
	Time    time.Time `json:"time,omitempty"`
	Message string    `json:"message,omitempty"`
}

type ErrorResponse struct {
	ErrorType string `json:"error_type"`
	Message   string `json:"message,omitempty"`
}
