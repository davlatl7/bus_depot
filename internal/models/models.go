package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `swaggerignore:"true"`
	FullName   string `json:"full_name"`
	Email      string `json:"email" gorm:"unique"`
	Phone      string `json:"phone"`
	Password   string `json:"-"`
	Role       string `json:"role"` // director, driver, mechanic

}

type Bus struct {
	gorm.Model `swaggerignore:"true"`
	Model1     string `json:"model1"`
	Number     string `json:"number"`
	Year       int    `json:"year"`
	Status     string `json:"status"` // в рейсе, в ремонте, свободен
	DriverID   *uint  `json:"driver_id"`
	MechanicID *uint  `json:"mechanic_id"`
}

type WorkSchedule struct {
	gorm.Model `swaggerignore:"true"`
	DriverID   uint   `json:"driver_id"`
	Driver     User   `json:"Driver" gorm:"foreignKey:DriverID;references:ID"`
	BusID      uint   `json:"bus_id"`
	Bus        Bus    `json:"Bus" gorm:"foreignKey:BusID;references:ID"`
	TimeRange  string `json:"time_range"`
	LineName   string `json:"line_name"`
}

type MyScheduleResponse struct {
	ID        uint         `json:"id"`
	TimeRange string       `json:"time_range"`
	LineName  string       `json:"line_name"`
	Bus       BusShortInfo `json:"bus"`
}

type BusShortInfo struct {
	Model1 string `json:"model1"`
	Number string `json:"number"`
	Status string `json:"status"`
}

type Report struct {
	gorm.Model `swaggerignore:"true"`
	MechanicID uint   `json:"mechanic_id"`
	Mechanic   User   `json:"mechanic" gorm:"foreignKey:MechanicID"`
	BusID      uint   `json:"bus_id"`
	Bus        Bus    `json:"bus" gorm:"foreignKey:BusID"`
	Status     string `json:"status"`
	Comment    string `json:"comment"`
}
