package base

import "time"

type BasicResponse struct {
	Time    time.Time `json:"time,omitempty"`
	Message string    `json:"message,omitempty"`
}
