package service

import (
	"bus_depot/internal/models"
	"bus_depot/internal/repository"
)

type BusService struct {
	repo *repository.BusRepository
}

func NewBusService(repo *repository.BusRepository) *BusService {
	return &BusService{repo: repo}
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
	bus, err := s.repo.GetByID(int(busID))
	if err != nil {
		return err
	}

	bus.DriverID = &driverID
	return s.repo.Update(bus)
}

func (s *BusService) AssignMechanic(busID, mechanicID uint) error {
	bus, err := s.repo.GetByID(int(busID))
	if err != nil {
		return err
	}

	bus.MechanicID = &mechanicID
	return s.repo.Update(bus)
}







