
// @title Bus Depot API
// @version 1.0
// @description API для управления автобусным парком
// @contact.name Поддержка
// @contact.email bus_depot@example.com
// @host localhost:8080
// @BasePath /
// @schemes http
package main


import (
	"bus_depot/internal/configs"
	"bus_depot/internal/handlers"
	"bus_depot/internal/logger"
	"bus_depot/internal/migrations"
	"bus_depot/internal/repository"
	"bus_depot/internal/service"
	"bus_depot/pkg/db"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "bus_depot/docs"

)

func main() {
	// ✅ Конфигурации
	if err := configs.ReadSettings(); err != nil {
		panic("Ошибка загрузки конфигурации: " + err.Error())
	}
	if err := logger.Init(); err != nil {
		panic("Ошибка инициализации логгера: " + err.Error())
	}

	// ✅ БД
	dbConn, err := db.InitDB()
	if err != nil {
		logger.Error.Fatalf("Ошибка инициализации БД: %v", err)
	}

	// ✅ Миграции
	if err := migrations.InitMigrations(dbConn); err != nil {
		logger.Error.Fatalf("Ошибка миграции: %v", err)
	}
	logger.Info.Println("Миграции завершены")

	// ✅ Репозитории и сервисы
	busRepo := repository.NewBusRepository(dbConn)
	userRepo := repository.NewUserRepository(dbConn)
	workScheduleRepo := repository.NewWorkScheduleRepository(dbConn)
	reportRepo := repository.NewReportRepository(dbConn)
	busService := service.NewBusService(busRepo)
	userService := service.NewUserService(userRepo)
	workScheduleService := service.NewWorkScheduleService(workScheduleRepo)
	reportService := service.NewReportService(reportRepo)
	authService := service.NewAuthService()

	// ✅ Хендлеры
	busHandler := handlers.NewBusHandler(busService)
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService, userService)
	workScheduleHandler := handlers.NewWorkScheduleHandler(workScheduleService)
	reportHandler := handlers.NewReportHandler(reportService)

	// ✅ Роутер
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))




	// ✅ Роуты
	handlers.InitRoutes(router, authHandler, userHandler, busHandler, workScheduleHandler, reportHandler)

	// ✅ Запуск сервера
	port := "8080"
	logger.Info.Printf("🚀 Сервер запущен на порту %s", port)
	if err := router.Run(":" + port); err != nil {
		logger.Error.Fatalf("❌ Ошибка запуска сервера: %v", err)
	}
}
