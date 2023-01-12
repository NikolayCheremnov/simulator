package task

type StepType string

const (
	CPU_STEP  = "cpu_step"
	IO_STEP   = "io_step"
	WAIT_STEP = "waiting"
)

// Task executable
type Task struct {
	PID              string     `json:"pid"` // pid of task
	RequiredSteps    []StepType `json:"steps"`
	NextRequiredStep int        `json:"next_required_step"`
	LiveTimeSteps    []StepType `json:"liveTimeSteps"` // real steps of live
	NextLiveTimeStep int        `json:"nex_live_time_step"`
}

func (task *Task) Step(isActiveOnCpu bool) {
	if task.IsCompleted() {
		panic("Try to make step for completed task")
	}

	if isActiveOnCpu {
		if task.NextStep() != CPU_STEP {
			panic("Scheduler error: choose not ready for CPU task")
		} else {
			task.LiveTimeSteps = append(task.LiveTimeSteps, CPU_STEP)
			task.NextRequiredStep++
			task.NextLiveTimeStep++
		}
	} else {
		if task.NextStep() == IO_STEP {
			task.LiveTimeSteps = append(task.LiveTimeSteps, IO_STEP)
			task.NextRequiredStep++
		} else {
			task.LiveTimeSteps = append(task.LiveTimeSteps, WAIT_STEP)
		}
		task.NextLiveTimeStep++
	}
}

func (task *Task) ReadyWorkOnCpu() bool {
	return task.NextStep() == CPU_STEP
}

func (task *Task) NextStep() StepType {
	return task.RequiredSteps[task.NextRequiredStep]
}

func (task *Task) IsCompleted() bool {
	return task.WorkLeft() == 0
}

func (task *Task) WorkLeft() int {
	return len(task.RequiredSteps) - task.NextRequiredStep
}
