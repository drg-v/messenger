package repository

import (
	"gorm.io/gorm"
	"messenger/pkg/entity"
)

type MessageRepository interface {
	Get(id int) (entity.Message, error)
	GetAllByUser(userId int) ([]entity.Message, error)
	Save(message entity.Message) error
	Update(message entity.Message) error
	Delete(id int) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (messageRepo *messageRepository) Get(id int) (entity.Message, error) {
	var message entity.Message
	err := messageRepo.db.First(&message, id).Error
	return message, err
}

func (messageRepo *messageRepository) GetAllByUser(id int) ([]entity.Message, error) {
	var messages []entity.Message
	err := messageRepo.db.Where("receiver = ?", id).Find(&messages).Error
	return messages, err
}

func (messageRepo *messageRepository) Save(message entity.Message) error {
	return messageRepo.db.Create(&message).Error
}

func (messageRepo *messageRepository) Update(message entity.Message) error {
	return messageRepo.db.Save(&message).Error
}

func (messageRepo *messageRepository) Delete(id int) error {
	return messageRepo.db.Delete(&entity.Message{}, id).Error
}
