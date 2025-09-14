package usecase

import (
	"todo-app/internal/domain"
)

// taskUsecase реализует интерфейс domain.TaskUsecase
type taskUsecase struct {
	taskService domain.TaskService
}

// NewTaskUsecase создает новый экземпляр usecase для задач
func NewTaskUsecase(taskService domain.TaskService) domain.TaskUsecase {
	return &taskUsecase{
		taskService: taskService,
	}
}

// AddTask добавляет новую задачу
func (u *taskUsecase) AddTask(title string) (*domain.Task, error) {
	return u.taskService.CreateTask(title)
}

// ListTasks возвращает список задач согласно фильтру
func (u *taskUsecase) ListTasks(filter domain.TaskFilter) ([]domain.Task, error) {
	return u.taskService.GetFilteredTasks(filter)
}

// CompleteTask отмечает задачу как выполненную
func (u *taskUsecase) CompleteTask(id uint) (*domain.Task, error) {
	task, err := u.taskService.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	// Если задача уже выполнена, возвращаем как есть
	if task.IsCompleted {
		return task, nil
	}

	return u.taskService.ToggleTaskCompletion(id)
}

// UncompleteTask отмечает задачу как невыполненную
func (u *taskUsecase) UncompleteTask(id uint) (*domain.Task, error) {
	task, err := u.taskService.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	// Если задача уже не выполнена, возвращаем как есть
	if !task.IsCompleted {
		return task, nil
	}

	return u.taskService.ToggleTaskCompletion(id)
}

// RemoveTask удаляет задачу
func (u *taskUsecase) RemoveTask(id uint) error {
	return u.taskService.DeleteTask(id)
}

// GetActiveTasks возвращает только активные (невыполненные) задачи
func (u *taskUsecase) GetActiveTasks() ([]domain.Task, error) {
	filter := domain.TaskFilter{
		Status: "active",
		SortBy: "created_at",
	}
	return u.taskService.GetFilteredTasks(filter)
}

// GetCompletedTasks возвращает только выполненные задачи
func (u *taskUsecase) GetCompletedTasks() ([]domain.Task, error) {
	filter := domain.TaskFilter{
		Status: "completed",
		SortBy: "created_at",
	}
	return u.taskService.GetFilteredTasks(filter)
}

// GetTaskByID возвращает задачу по ID
func (u *taskUsecase) GetTaskByID(id uint) (*domain.Task, error) {
	return u.taskService.GetTaskByID(id)
}
