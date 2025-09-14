package domain

import (
	"time"
)

// Task представляет задачу в системе
type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	IsCompleted bool      `json:"is_completed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// TaskFilter представляет фильтры для задач
type TaskFilter struct {
	Status string `json:"status"`  // "all", "active", "completed"
	SortBy string `json:"sort_by"` // "created_at"
}

// TaskRepository интерфейс для работы с задачами в репозитории
type TaskRepository interface {
	Create(task *Task) error
	GetAll() ([]Task, error)
	GetByID(id uint) (*Task, error)
	Update(task *Task) error
	Delete(id uint) error
	GetByFilter(filter TaskFilter) ([]Task, error)
}

// TaskService интерфейс для бизнес-логики задач
type TaskService interface {
	CreateTask(title string) (*Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id uint) (*Task, error)
	ToggleTaskCompletion(id uint) (*Task, error)
	DeleteTask(id uint) error
	GetFilteredTasks(filter TaskFilter) ([]Task, error)
	ValidateTaskTitle(title string) error
}

// TaskUsecase интерфейс для слоя использования
type TaskUsecase interface {
	AddTask(title string) (*Task, error)
	ListTasks(filter TaskFilter) ([]Task, error)
	CompleteTask(id uint) (*Task, error)
	UncompleteTask(id uint) (*Task, error)
	RemoveTask(id uint) error
	GetActiveTasks() ([]Task, error)
	GetCompletedTasks() ([]Task, error)
	GetTaskByID(id uint) (*Task, error)
}
