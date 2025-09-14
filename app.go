package main

import (
	"context"
	"todo-app/internal/handler"
)

// App struct
type App struct {
	ctx         context.Context
	taskHandler *handler.TaskHandler
}

// NewApp creates a new App application struct
func NewApp(taskHandler *handler.TaskHandler) *App {
	return &App{
		taskHandler: taskHandler,
	}
}

// startup is called when the app starts up
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name (можно оставить для тестирования)
func (a *App) Greet(name string) string {
	return "Hello " + name + ", It's show time!"
}

// TaskHandler methods - эти методы будут доступны из frontend

// AddTask добавляет новую задачу
func (a *App) AddTask(title string) (interface{}, error) {
	return a.taskHandler.AddTask(a.ctx, title)
}

// GetAllTasks возвращает все задачи
func (a *App) GetAllTasks() (interface{}, error) {
	return a.taskHandler.GetAllTasks(a.ctx)
}

// GetActiveTasks возвращает активные задачи
func (a *App) GetActiveTasks() (interface{}, error) {
	return a.taskHandler.GetActiveTasks(a.ctx)
}

// GetCompletedTasks возвращает выполненные задачи
func (a *App) GetCompletedTasks() (interface{}, error) {
	return a.taskHandler.GetCompletedTasks(a.ctx)
}

// ToggleTask переключает статус задачи
func (a *App) ToggleTask(id uint) (interface{}, error) {
	// Сначала получаем задачу, чтобы узнать её текущий статус
	task, err := a.taskHandler.GetTaskByID(a.ctx, id)
	if err != nil {
		return nil, err
	}

	// Переключаем статус
	if task.IsCompleted {
		return a.taskHandler.UncompleteTask(a.ctx, id)
	} else {
		return a.taskHandler.CompleteTask(a.ctx, id)
	}
}

// DeleteTask удаляет задачу
func (a *App) DeleteTask(id uint) error {
	return a.taskHandler.DeleteTask(a.ctx, id)
}

// GetFilteredTasks возвращает отфильтрованные задачи
func (a *App) GetFilteredTasks(status string) (interface{}, error) {
	return a.taskHandler.GetFilteredTasks(a.ctx, status)
}
