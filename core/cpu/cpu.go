package cpu

type Status string

const (
	WORKING = "work"
	IDLE    = "downtime"
)

type WorkingInfo struct {
	StatusCache []Status `json:"status_cache"`
}

func (workingInfo *WorkingInfo) CalculateCpuActivityMetrics() float64 {
	working := 0
	for _, status := range workingInfo.StatusCache {
		if status == WORKING {
			working++
		}
	}
	return float64(working) / float64(len(workingInfo.StatusCache)) * 100.0
}
