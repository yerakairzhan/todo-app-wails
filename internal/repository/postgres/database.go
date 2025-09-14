package postgres

import (
	"todo-app/internal/config"
	"todo-app/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDatabase инициализирует подключение к PostgreSQL и выполняет миграции
func InitDatabase(cfg *config.Config) (*gorm.DB, error) {
	// Подключение к базе данных
	db, err := gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Автоматическая миграция схемы
	err = db.AutoMigrate(&domain.Task{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
