package service

import (
	"bus_depot/internal/models"
	"bus_depot/internal/repository"
)

type ReportService struct {
	repo *repository.ReportRepository
}

func NewReportService(r *repository.ReportRepository) *ReportService {
	return &ReportService{r}
}

func (s *ReportService) CreateReport(report *models.Report) error {
	return s.repo.CreateReport(report)
}

func (s *ReportService) GetAllReports() ([]models.Report, error) {
	return s.repo.GetAllReports()
}

func (s *ReportService) GetReportByID(id uint) (*models.Report, error) {
	return s.repo.GetReportByID(id)
}

func (s *ReportService) DeleteReport(id uint) error {
	return s.repo.DeleteReport(id)
}