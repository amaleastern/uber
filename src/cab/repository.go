package cab

import (
	"uber/src/models"

	"gorm.io/gorm"
)

// Repository interface
type Repository interface {
	GetCabsByLatLong(latitude, longitude float64) ([]models.Cab, error)
}

// RepositoryImpl struct
type RepositoryImpl struct {
	DB *gorm.DB
}

// NewRepository func
func NewRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{DB: db}
}

// GetCabsByLatLong Get cabs by Latitude and Longitude
func (r *RepositoryImpl) GetCabsByLatLong(latitude, longitude float64) ([]models.Cab, error) {
	radius := .1
	var cabs []models.Cab

	err := r.DB.Joins("JOIN location ON location.id = cab.id_location").
		Where("(location.latitude BETWEEN ? AND ?) AND (location.longitude BETWEEN ? AND ?) AND active = ?",
			latitude-radius, latitude+radius, longitude-radius, longitude+radius, true).Find(&cabs).Error
	return cabs, err
}
