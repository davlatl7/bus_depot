// internal/service/work_schedule_service.go
package service

import (
	"bus_depot/internal/models"
	"bus_depot/internal/repository"
	"errors"
)

type WorkScheduleService struct {
	repo *repository.WorkScheduleRepository
}

func NewWorkScheduleService(r *repository.WorkScheduleRepository) *WorkScheduleService {
	return &WorkScheduleService{r}
}

func (s *WorkScheduleService) CreateSchedule(schedule *models.WorkSchedule) error {
	driverBusy, err := s.repo.IsDriverBusy(schedule.DriverID)
	if err != nil {
		return err
	}
	if driverBusy {
		return errors.New("у этого водителя уже есть график")
	}

	busBusy, err := s.repo.IsBusBusy(schedule.BusID)
	if err != nil {
		return err
	}
	if busBusy {
		return errors.New("этот автобус уже используется в другом графике")
	}

	return s.repo.CreateSchedule(schedule)
}

func (s *WorkScheduleService) GetAllSchedules() ([]models.WorkSchedule, error) {
	return s.repo.GetAllSchedules()
}

func (s *WorkScheduleService) GetScheduleByID(id uint) (*models.WorkSchedule, error) {
	return s.repo.GetScheduleByID(id)
}

func (s *WorkScheduleService) UpdateSchedule(schedule *models.WorkSchedule) error {
	return s.repo.UpdateSchedule(schedule)
}

func (s *WorkScheduleService) DeleteSchedule(id uint) error {
	return s.repo.DeleteSchedule(id)
}


func (s *WorkScheduleService) GetByDriverID(driverID uint) ([]models.WorkSchedule, error) {
	return s.repo.GetByDriverID(driverID)
}