// main.go
package main

import (
	"embed"
	"fmt"
	"log"

	"todo-app/internal/config"
	"todo-app/internal/handler"
	"todo-app/internal/repository/postgres"
	"todo-app/internal/service"
	"todo-app/internal/usecase"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Загружаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Инициализируем базу данных
	db, err := postgres.InitDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Создаем слои архитектуры
	taskRepo := postgres.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskUsecase := usecase.NewTaskUsecase(taskService)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	// Create an instance of the app structure
	app := NewApp(taskHandler)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Todo App",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Fullscreen:       false,
		WindowStartState: options.Normal,
		Bind: []interface{}{
			app, // Bind the app instance to generate JavaScript bindings
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func init() {
	// Проверяем подключение к базе данных при запуске
	fmt.Println("Initializing Todo App...")
}
