package services

import (
	"golang_api/dto"
	"golang_api/models"
	"golang_api/repository"

	"golang.org/x/crypto/bcrypt"
)

//AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyAccount(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) models.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (service *authService) VerifyAccount(email string, password string) interface{} {
	res := service.userRepository.VerifyAccount(email, password)
	if user, ok := res.(models.User); ok {
		errorCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if errorCompare == nil {
			return user
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) models.User {
	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	res := service.userRepository.CreateUser(newUser)
	return res
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}
