package cpu

type Status string

const (
	WORKING = "work"
	IDLE    = "downtime"
)

type WorkingInfo struct {
	StatusCache []Status `json:"status_cache"`
}
