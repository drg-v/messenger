package repository

import (
	"gorm.io/gorm"
	"messenger/pkg/entity"
)

type UserRepository interface {
	Get(id int) (entity.User, error)
	GetAll() ([]entity.User, error)
	Save(user entity.User) error
	Update(user entity.User) error
	Delete(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (userRepo *userRepository) Get(id int) (entity.User, error) {
	var user entity.User
	err := userRepo.db.First(&user, id).Error
	return user, err
}

func (userRepo *userRepository) GetAll() ([]entity.User, error) {
	var users []entity.User
	err := userRepo.db.Order("id").Find(&users).Error
	return users, err
}

func (userRepo *userRepository) Save(user entity.User) error {
	return userRepo.db.Create(&user).Error
}

func (userRepo *userRepository) Update(user entity.User) error {
	return userRepo.db.Save(&user).Error
}

func (userRepo *userRepository) Delete(id int) error {
	return userRepo.db.Delete(&entity.User{}, id).Error
}
