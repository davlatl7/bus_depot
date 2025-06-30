package repository

import (
	"bus_depot/internal/models"
	"gorm.io/gorm"
)

type WorkScheduleRepository struct {
	db *gorm.DB
}

func NewWorkScheduleRepository(db *gorm.DB) *WorkScheduleRepository {
	return &WorkScheduleRepository{db}
}

func (r *WorkScheduleRepository) CreateSchedule(schedule *models.WorkSchedule) error {
	return r.db.Create(schedule).Error
}

func (r *WorkScheduleRepository) GetAllSchedules() ([]models.WorkSchedule, error) {
	var schedules []models.WorkSchedule
	err := r.db.Debug().Preload("Driver").Preload("Bus").Find(&schedules).Error
	return schedules, err
}

func (r *WorkScheduleRepository) GetScheduleByID(id uint) (*models.WorkSchedule, error) {
	var schedule models.WorkSchedule
	err := r.db.Preload("Driver").Preload("Bus").First(&schedule, id).Error
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *WorkScheduleRepository) UpdateSchedule(schedule *models.WorkSchedule) error {
	return r.db.Save(schedule).Error
}

func (r *WorkScheduleRepository) DeleteSchedule(id uint) error {
	return r.db.Delete(&models.WorkSchedule{}, id).Error
}

func (r *WorkScheduleRepository) GetByDriverID(driverID uint) ([]models.WorkSchedule, error) {
	var schedules []models.WorkSchedule
	err := r.db.Preload("Bus").Where("driver_id = ?", driverID).Find(&schedules).Error
	return schedules, err
}

func (r *WorkScheduleRepository) IsDriverBusy(driverID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.WorkSchedule{}).
		Where("driver_id = ?", driverID).
		Count(&count).Error
	return count > 0, err
}

func (r *WorkScheduleRepository) IsBusBusy(busID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.WorkSchedule{}).
		Where("bus_id = ?", busID).
		Count(&count).Error
	return count > 0, err
}

