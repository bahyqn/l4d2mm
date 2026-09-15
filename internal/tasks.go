package internal

type TaskType int
type TaskStatus int

const (
	TaskNone TaskType = iota
	TaskTypeCollection
	TaskTypeDelete
	TaskTypeCRC
)
const (
	TaskStatusIdle TaskStatus = iota
	TaskStatusRunning
	TaskStatusCompleted
	TaskStatusFailed
)

type History struct {
	Id     string
	Type   TaskType
	Status TaskStatus
}

type Task struct {
	TaskType TaskType
	Status   TaskStatus
	Selected map[string]TaskStatus
	History  []History
}

func NewTask() *Task {
	return &Task{
		TaskType: TaskNone,
		Status:   TaskStatusIdle,
		Selected: map[string]TaskStatus{},
	}
}

func (task *Task) AddTask(modId string) {
	if _, ok := task.Selected[modId]; ok {
		return
	}

	task.Selected[modId] = TaskStatusIdle
}
func (task *Task) RemoveTask(modId string) {
	if _, ok := task.Selected[modId]; ok {
		return
	}

	delete(task.Selected, modId)
}

func (task *Task) IsSelected(modId string) bool {
	if _, ok := task.Selected[modId]; !ok {
		return false
	}
	return true
}

func (task *Task) SetupMode(mode TaskType) {
	task.TaskType = mode
}

func (task *Task) SwitchTaskStatus(status TaskStatus) {
	task.Status = status
}

func (task *Task) SwitchSelectedTaskStatus(modId string, status TaskStatus) {
	if _, ok := task.Selected[modId]; !ok {
		return
	}

	task.Selected[modId] = status
}

func (task *Task) Cancel() {
	clear(task.Selected)
}
