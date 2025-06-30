package repository

import (
	"bus_depot/internal/models"
	"gorm.io/gorm"
)

type ReportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{db}
}

func (r *ReportRepository) CreateReport(report *models.Report) error {
	return r.db.Create(report).Error
}

func (r *ReportRepository) GetAllReports() ([]models.Report, error) {
	var reports []models.Report
	err := r.db.Preload("Mechanic").Preload("Bus").Find(&reports).Error
	return reports, err
}

func (r *ReportRepository) GetReportByID(id uint) (*models.Report, error) {
	var report models.Report
	err := r.db.Preload("Mechanic").Preload("Bus").First(&report, id).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *ReportRepository) DeleteReport(id uint) error {
	return r.db.Delete(&models.Report{}, id).Error
}