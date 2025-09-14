package service

import (
	"errors"
	"strings"
	"todo-app/internal/domain"
)

// taskService реализует интерфейс domain.TaskService
type taskService struct {
	taskRepo domain.TaskRepository
}

// NewTaskService создает новый экземпляр сервиса задач
func NewTaskService(taskRepo domain.TaskRepository) domain.TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

// CreateTask создает новую задачу после валидации
func (s *taskService) CreateTask(title string) (*domain.Task, error) {
	// Валидация заголовка
	if err := s.ValidateTaskTitle(title); err != nil {
		return nil, err
	}

	// Создание задачи
	task := &domain.Task{
		Title:       strings.TrimSpace(title),
		IsCompleted: false,
	}

	err := s.taskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// GetAllTasks возвращает все задачи
func (s *taskService) GetAllTasks() ([]domain.Task, error) {
	return s.taskRepo.GetAll()
}

// GetTaskByID возвращает задачу по ID
func (s *taskService) GetTaskByID(id uint) (*domain.Task, error) {
	if id == 0 {
		return nil, errors.New("invalid task ID")
	}
	return s.taskRepo.GetByID(id)
}

// ToggleTaskCompletion переключает статус выполнения задачи
func (s *taskService) ToggleTaskCompletion(id uint) (*domain.Task, error) {
	task, err := s.taskRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Переключаем статус
	task.IsCompleted = !task.IsCompleted

	err = s.taskRepo.Update(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// DeleteTask удаляет задачу
func (s *taskService) DeleteTask(id uint) error {
	if id == 0 {
		return errors.New("invalid task ID")
	}

	// Проверяем существование задачи
	_, err := s.taskRepo.GetByID(id)
	if err != nil {
		return errors.New("task not found")
	}

	return s.taskRepo.Delete(id)
}

// GetFilteredTasks возвращает отфильтрованные задачи
func (s *taskService) GetFilteredTasks(filter domain.TaskFilter) ([]domain.Task, error) {
	return s.taskRepo.GetByFilter(filter)
}

// ValidateTaskTitle валидирует заголовок задачи
func (s *taskService) ValidateTaskTitle(title string) error {
	title = strings.TrimSpace(title)

	if title == "" {
		return errors.New("task title cannot be empty")
	}

	if len(title) > 255 {
		return errors.New("task title is too long (max 255 characters)")
	}

	return nil
}
