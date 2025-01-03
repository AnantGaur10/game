package repositories

import (
	"game/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := repo.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (repo *UserRepository) GetUserByID(ID uint) (*models.User, error) {
	var user models.User
	result := repo.db.Where("id = ?", ID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
