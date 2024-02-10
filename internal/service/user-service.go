package service

import (
	"fmt"
	"project/internal/dto"
	"project/internal/helper"
	"project/internal/models"
	"project/internal/repository"

	"github.com/mashingan/smapping"
)

// UserService is a contract.....
type UserService interface {
	Update(user dto.UserUpdateDTO) (*models.User, error)
	Profile(userID string) models.User
	ListUserByFilter(dtoList dto.ListUserDTO) (*[]models.User, *int64, error)
}

type userService struct {
	userRepository repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) (*models.User, error) {
	userToUpdate := models.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		helper.WriteError(err.Error())
		return nil, err
	}
	userToUpdate.ID = user.ID
	updatedUser, err := service.userRepository.UpdateUser(userToUpdate)
	if err != nil {
		helper.WriteError(err.Error())
		fmt.Println(err.Error())
		return nil, err
	}
	return updatedUser, nil
}

func (service *userService) Profile(userID string) models.User {
	return service.userRepository.ProfileUser(userID)
}
func (service *userService) ListUserByFilter(dtoList dto.ListUserDTO) (*[]models.User, *int64, error) {
	return service.userRepository.ListUserByFilter(dtoList)
}
