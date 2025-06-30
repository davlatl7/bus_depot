package migrations

import (
	"log"
	"gorm.io/gorm"
	"bus_depot/internal/models"
	
)


func InitMigrations(db *gorm.DB) error {
	
	if err := db.AutoMigrate(&models.User{}, &models.Bus{}, &models.WorkSchedule{}, &models.Report{}); err != nil {
		log.Printf("Ошибка миграции: %v", err)
		return err
	}
	log.Println("Миграция таблиц выполнена")


	return nil
}
