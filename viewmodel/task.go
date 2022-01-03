package viewmodel

import (
	"go-ws/task-app/models"
)

type TasksViewModel struct {
	Tasks []models.Task
}

func MakeTaskViewModel(tasks []models.Task) *TasksViewModel {
	return &TasksViewModel{Tasks: tasks}
}
