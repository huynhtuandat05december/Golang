package services

import (
	"golang_api/dto"
	"golang_api/models"
	"golang_api/repository"
)

type UserService interface {
	GetProfileUser(userID uint64) models.User
	UpdateProfileUser(userUpdateDTO dto.UserUpdateDTO) models.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (service *userService) GetProfileUser(userID uint64) models.User {
	return service.userRepository.ProfileUser(userID)
}

func (service *userService) UpdateProfileUser(userUpdateDTO dto.UserUpdateDTO) models.User {
	var newProfileUser models.User
	if userUpdateDTO.Email != "" {
		newProfileUser.Email = userUpdateDTO.Email
	}
	if userUpdateDTO.Name != "" {
		newProfileUser.Email = userUpdateDTO.Name
	}
	if userUpdateDTO.Password != "" {
		newProfileUser.Password = userUpdateDTO.Password
	}
	return service.userRepository.UpdateUser(newProfileUser)

}
