package user

import (
	"uber/src/models"

	"gorm.io/gorm"
)

// Repository interface
type Repository interface {
	Getuser(username string) (*models.User, error)
	CreateUser(user *models.User) error
}

// RepositoryImpl struct
type RepositoryImpl struct {
	DB *gorm.DB
}

// NewRepository func
func NewRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{DB: db}
}

// Getuser Get user by Username
func (r *RepositoryImpl) Getuser(username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

// CreateUser
func (r *RepositoryImpl) CreateUser(user *models.User) error {
	return r.DB.Create(&user).Error
}
