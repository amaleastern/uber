package cab

import (
	"strconv"
	"uber/src/models"
)

// Service interface
type Service interface {
	GetNearbyCabs(latitude, longitude string) ([]models.Cab, error)
}

// ServiceImpl struct
type ServiceImpl struct {
	Repository Repository
}

// NewService func
func NewService(repository Repository) Service {
	return &ServiceImpl{Repository: repository}
}

// GetNearbyCabs Get cabs by Latitude and Longitude
func (s *ServiceImpl) GetNearbyCabs(latitude, longitude string) ([]models.Cab, error) {
	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return nil, err
	}

	long, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return nil, err
	}

	return s.Repository.GetCabsByLatLong(lat, long)
}
