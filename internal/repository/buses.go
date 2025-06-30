package repository

import (
	"bus_depot/internal/models"
	"gorm.io/gorm"
)

type BusRepository struct {
	db *gorm.DB
}

func NewBusRepository(db *gorm.DB) *BusRepository {
	return &BusRepository{db: db}
}

func (r *BusRepository) Create(bus *models.Bus) error {
	return r.db.Create(bus).Error
}

func (r *BusRepository) GetAll() ([]models.Bus, error) {
	var buses []models.Bus
	err := r.db.Find(&buses).Error
	return buses, err
}

func (r *BusRepository) GetByID(id int) (*models.Bus, error) {
	var bus models.Bus
	err := r.db.First(&bus, id).Error
	if err != nil {
		return nil, err
	}
	return &bus, nil
}


func (r *BusRepository) Update(bus *models.Bus) error {
	return r.db.Save(bus).Error
}

func (r *BusRepository) Delete(id int) error {
	return r.db.Delete(&models.Bus{}, id).Error
}

func (r *BusRepository) AssignDriver(busID uint, driverID uint) error {
	return r.db.Model(&models.Bus{}).
		Where("id = ?", busID).
		Update("driver_id", driverID).Error
}

func (r *BusRepository) AssignMaster(busID uint, masterID uint) error {
	return r.db.Model(&models.Bus{}).
		Where("id = ?", busID).
		Update("master_id", masterID).Error
}

func (r *BusRepository) GetByDriverID(driverID uint) (*models.Bus, error) {
	var bus models.Bus
	if err := r.db.Where("driver_id = ?", driverID).First(&bus).Error; err != nil {
		return nil, err
	}
	return &bus, nil
}
