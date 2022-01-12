package taskmanager



type CreateTaskDto struct {
	Title string `json:"title"`
	Description string `json:"description"`
}


type UpdateTaskDto struct {
	TaskId string `json:"taskId"`
	Title string `json:"title"`
	Description string `json:"description"`
}


type GetTasksByIdDto struct{
	TaskId string `uri:"taskId"`
}


type TaskDetailsDto struct {
	TaskId uint `json:"taskId"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type ApproveTaskDto struct {
	TaskId uint `json:"taskId"`
}