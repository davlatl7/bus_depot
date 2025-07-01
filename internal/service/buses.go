package service

import (
	"bus_depot/internal/models"
	"bus_depot/internal/repository"
	"fmt"
)

type BusService struct {
	repo *repository.BusRepository
	userRepo  *repository.UserRepository
}

func NewBusService(repo *repository.BusRepository,  userRepo *repository.UserRepository) *BusService {
	 return &BusService{
        repo:     repo,
        userRepo: userRepo,
    }
}

func (s *BusService) CreateBus(bus *models.Bus) error {
	return s.repo.Create(bus)
}

func (s *BusService) GetAllBuses() ([]models.Bus, error) {
	return s.repo.GetAll()
}

func (s *BusService) GetBusByID(id int) (*models.Bus, error) {
	return s.repo.GetByID(id)
}

func (s *BusService) UpdateBus(bus *models.Bus) error {
	return s.repo.Update(bus)
}

func (s *BusService) DeleteBus(id int) error {
	return s.repo.Delete(id)
}

func (s *BusService) AssignDriver(busID, driverID uint) error {
	exists, err := s.userRepo.ExistsByID(driverID)
	if err != nil {
		return fmt.Errorf("ошибка при проверке водителя: %v", err)
	}
	if !exists {
		return fmt.Errorf("водитель с ID %d не найден", driverID)
	}

	
	existingBus, err := s.repo.GetByDriverID(driverID)
	if err == nil && existingBus != nil {
		return fmt.Errorf("водитель уже назначен на автобус с ID %d", existingBus.ID)
	}

	return s.repo.AssignDriver(busID, driverID)
}


func (s *BusService) AssignMechanic(busID, mechanicID uint) error {
	// Проверка: механик существует?
	exists, err := s.userRepo.ExistsByID(mechanicID)
	if err != nil {
		return fmt.Errorf("ошибка при проверке механика: %v", err)
	}
	if !exists {
		return fmt.Errorf("механик с ID %d не найден", mechanicID)
	}

	
	existingBus, err := s.repo.GetByMechanicID(mechanicID)
	if err == nil && existingBus != nil {
		return fmt.Errorf("механик уже назначен на автобус с ID %d", existingBus.ID)
	}

	return s.repo.AssignMaster(busID, mechanicID)
}

