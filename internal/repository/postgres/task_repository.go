package postgres

import (
	"todo-app/internal/domain"

	"gorm.io/gorm"
)

// taskRepository реализует интерфейс domain.TaskRepository
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository создает новый экземпляр репозитория задач
func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepository{
		db: db,
	}
}

// Create создает новую задачу в базе данных
func (r *taskRepository) Create(task *domain.Task) error {
	return r.db.Create(task).Error
}

// GetAll возвращает все задачи из базы данных
func (r *taskRepository) GetAll() ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.db.Order("created_at DESC").Find(&tasks).Error
	return tasks, err
}

// GetByID возвращает задачу по ID
func (r *taskRepository) GetByID(id uint) (*domain.Task, error) {
	var task domain.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// Update обновляет задачу в базе данных
func (r *taskRepository) Update(task *domain.Task) error {
	return r.db.Save(task).Error
}

// Delete удаляет задачу из базы данных
func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Task{}, id).Error
}

// GetByFilter возвращает задачи согласно фильтру
func (r *taskRepository) GetByFilter(filter domain.TaskFilter) ([]domain.Task, error) {
	var tasks []domain.Task
	query := r.db.Model(&domain.Task{})

	// Фильтрация по статусу
	switch filter.Status {
	case "active":
		query = query.Where("is_completed = ?", false)
	case "completed":
		query = query.Where("is_completed = ?", true)
	case "all":
		// Не добавляем условие - показываем все
	}

	// Сортировка
	switch filter.SortBy {
	case "created_at":
		query = query.Order("created_at DESC")
	default:
		query = query.Order("created_at DESC")
	}

	err := query.Find(&tasks).Error
	return tasks, err
}
