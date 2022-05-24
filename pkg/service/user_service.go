package service

import (
	"errors"
	"messenger/pkg/dto"
	"messenger/pkg/entity"
	"messenger/pkg/repository"
)

type UserService interface {
	Get(id int) (dto.UserDto, error)
	GetAll() ([]dto.UserDto, error)
	Create(user dto.UserDto) error
	Update(user dto.UserDto) error
	Delete(id int) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepository: userRepo}
}

func (userService *userService) Get(id int) (dto.UserDto, error) {
	userEntity, err := userService.userRepository.Get(id)
	if err != nil {
		return dto.UserDto{}, errors.New("user service - unable to find the user")
	}
	userDto := dto.UserDto{
		ID:       userEntity.ID,
		Name:     userEntity.Name,
		Password: userEntity.Password,
	}
	return userDto, nil
}

func (userService *userService) GetAll() ([]dto.UserDto, error) {
	userEntities, err := userService.userRepository.GetAll()
	if err != nil {
		return []dto.UserDto{}, errors.New("user service - unable to find all users")
	}
	userDtoSlice := make([]dto.UserDto, 0, len(userEntities))
	for _, val := range userEntities {
		userDtoSlice = append(userDtoSlice, dto.UserDto{
			ID:       val.ID,
			Name:     val.Name,
			Password: val.Password,
		})
	}
	return userDtoSlice, nil
}

func (userService *userService) Create(user dto.UserDto) error {
	bankEntity := entity.User{
		Name:     user.Name,
		Password: user.Password,
	}
	err := userService.userRepository.Save(bankEntity)
	if err != nil {
		return errors.New("user service - error creating new user")
	}
	return nil
}

func (userService *userService) Update(user dto.UserDto) error {
	userEntity := entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Password: user.Password,
	}
	err := userService.userRepository.Update(userEntity)
	if err != nil {
		return errors.New("user service - error updating user")
	}
	return nil
}

func (userService *userService) Delete(id int) error {
	err := userService.userRepository.Delete(id)
	if err != nil {
		return errors.New("user service - error deleting user")
	}
	return nil
}
