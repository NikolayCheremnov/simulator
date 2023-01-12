package task

type StepType string

const (
	CPU_STEP  = "cpu_tick"
	IO_STEP   = "io_step"
	WAIT_STEP = "io_step"
)

// Task executable
type Task struct {
	PID              string     `json:"pid"` // pid of task
	RequiredSteps    []StepType `json:"steps"`
	NextRequiredStep int        `json:"next_required_step"`
	LiveTimeSteps    []StepType `json:"liveTimeSteps"` // real steps of live
	NexLiveTimeStep  int        `json:"nex_live_time_step"`
}
