package service

import (
	"errors"
	"fmt"
	"messenger/pkg/dto"
	"messenger/pkg/entity"
	"messenger/pkg/repository"
)

type MessageService interface {
	Get(id int) (dto.MessageDto, error)
	GetAllByUser(id int) ([]dto.MessageDto, error)
	Create(message dto.MessageDto) error
	Update(message dto.MessageDto) error
	Delete(id int) error
}

type messageService struct {
	messageRepository repository.MessageRepository
}

func NewMessageService(messageRepo repository.MessageRepository) MessageService {
	return &messageService{messageRepository: messageRepo}
}

func (messageService *messageService) Get(id int) (dto.MessageDto, error) {
	messageEntity, err := messageService.messageRepository.Get(id)
	if err != nil {
		return dto.MessageDto{}, errors.New("message service - unable to find the message")
	}
	messageDto := dto.MessageDto{
		ID:       messageEntity.ID,
		Sender:   messageEntity.Sender,
		Receiver: messageEntity.Receiver,
		Text:     messageEntity.Text,
	}
	return messageDto, nil
}

func (messageService *messageService) GetAllByUser(id int) ([]dto.MessageDto, error) {
	messageEntities, err := messageService.messageRepository.GetAllByUser(id)
	fmt.Println(err)
	if err != nil {
		return []dto.MessageDto{}, errors.New("message service - unable to find messages by user")
	}
	messageDtoSlice := make([]dto.MessageDto, 0, len(messageEntities))
	for _, val := range messageEntities {
		messageDtoSlice = append(messageDtoSlice, dto.MessageDto{
			ID:       val.ID,
			Sender:   val.Sender,
			Receiver: val.Receiver,
			Text:     val.Text,
		})
	}
	return messageDtoSlice, nil
}

func (messageService *messageService) Create(message dto.MessageDto) error {
	messageEntity := entity.Message{
		Sender:   message.Sender,
		Receiver: message.Receiver,
		Text:     message.Text,
	}
	err := messageService.messageRepository.Save(messageEntity)
	if err != nil {
		return errors.New("message service - error creating new message")
	}
	return nil
}

func (messageService *messageService) Update(message dto.MessageDto) error {
	messageEntity := entity.Message{
		ID:       message.ID,
		Sender:   message.Sender,
		Receiver: message.Receiver,
		Text:     message.Text,
	}
	err := messageService.messageRepository.Update(messageEntity)
	if err != nil {
		return errors.New("message service - error updating message")
	}
	return nil
}

func (messageService *messageService) Delete(id int) error {
	err := messageService.messageRepository.Delete(id)
	if err != nil {
		return errors.New("message service - error deleting message")
	}
	return nil
}
