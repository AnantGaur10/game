package repositories

import(
	"gorm.io/gorm"
	"game/models"
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

func (repo *UserRepository) GetUserByEmail(email string) (*models.User,error) {
	var user models.User;
	result := repo.db.Where("email = ?",email).First(&user)
	if result.Error != nil {
		return nil,result.Error
	}
	return &user,nil
}
