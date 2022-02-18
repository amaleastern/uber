package user

import (
	"uber/src/models"
	"uber/src/util"
)

// Service interface
type Service interface {
	CreateUser(user *models.User) error
}

// ServiceImpl struct
type ServiceImpl struct {
	Repository Repository
}

// NewService func
func NewService(repository Repository) Service {
	return &ServiceImpl{Repository: repository}
}

// CreateUser
func (s *ServiceImpl) CreateUser(user *models.User) (err error) {
	user.Password, err = util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	return s.Repository.CreateUser(user)
}
