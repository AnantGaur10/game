package services

import (
	"log"

	"game/models"
	"game/pkg/auth"
	"game/repositories"
)

type UserService struct {
	Repository *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repository: repo}
}

func (s *UserService) CreateUser(user *models.User) (string, error) {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		log.Printf("Failed to hash the password for user %s: %v", user.Email, err)
		return "",err
	}
	user.Password = hashedPassword
	err = s.Repository.CreateUser(user)
    if err != nil {
        log.Printf("Failed to create user in the database: %v", err)
        return "", err
    }

	token,err := auth.GenerateJWT(int(user.ID))
	if err != nil {
		log.Printf("Failed to generate JWT for user ID %d: %v", user.ID, err)
		return "",err
	}
	return token,nil
}

func (s *UserService) SignUser(user *models.User) (string,error) {
	var email string = user.Email;
	user,err := s.Repository.GetUserByEmail(email)
	if err != nil {
		log.Println("Error while getting user by email {}",err)
		return "",err
	}
	token,err := auth.GenerateJWT(int(user.ID))
	if err != nil {
		log.Printf("Failed to generate JWT for user ID %d: %v", user.ID, err)
		return "",err
	}
	return token,nil

}