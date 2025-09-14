// internal/handler/task_handler.go
package handler

import (
	"context"
	"todo-app/internal/domain"
)

// TaskHandler обрабатывает запросы от frontend
type TaskHandler struct {
	taskUsecase domain.TaskUsecase
}

// NewTaskHandler создает новый экземпляр обработчика задач
func NewTaskHandler(taskUsecase domain.TaskUsecase) *TaskHandler {
	return &TaskHandler{
		taskUsecase: taskUsecase,
	}
}

// AddTask добавляет новую задачу (вызывается из frontend)
func (h *TaskHandler) AddTask(ctx context.Context, title string) (*domain.Task, error) {
	return h.taskUsecase.AddTask(title)
}

// GetAllTasks возвращает все задачи
func (h *TaskHandler) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	filter := domain.TaskFilter{
		Status: "all",
		SortBy: "created_at",
	}
	return h.taskUsecase.ListTasks(filter)
}

// GetActiveTasks возвращает только активные задачи
func (h *TaskHandler) GetActiveTasks(ctx context.Context) ([]domain.Task, error) {
	return h.taskUsecase.GetActiveTasks()
}

// GetCompletedTasks возвращает только выполненные задачи
func (h *TaskHandler) GetCompletedTasks(ctx context.Context) ([]domain.Task, error) {
	return h.taskUsecase.GetCompletedTasks()
}

// CompleteTask отмечает задачу как выполненную
func (h *TaskHandler) CompleteTask(ctx context.Context, id uint) (*domain.Task, error) {
	return h.taskUsecase.CompleteTask(id)
}

// UncompleteTask отмечает задачу как невыполненную
func (h *TaskHandler) UncompleteTask(ctx context.Context, id uint) (*domain.Task, error) {
	return h.taskUsecase.UncompleteTask(id)
}

// DeleteTask удаляет задачу
func (h *TaskHandler) DeleteTask(ctx context.Context, id uint) error {
	return h.taskUsecase.RemoveTask(id)
}

// GetFilteredTasks возвращает отфильтрованные задачи
func (h *TaskHandler) GetFilteredTasks(ctx context.Context, status string) ([]domain.Task, error) {
	filter := domain.TaskFilter{
		Status: status,
		SortBy: "created_at",
	}
	return h.taskUsecase.ListTasks(filter)
}

// GetTaskByID возвращает задачу по ID
func (h *TaskHandler) GetTaskByID(ctx context.Context, id uint) (*domain.Task, error) {
	return h.taskUsecase.GetTaskByID(id)
}
