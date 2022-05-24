package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"messenger/pkg/dto"
	"messenger/pkg/service"
	"net/http"
	"strconv"
)

type MessageHandler interface {
	CreateMessage() gin.HandlerFunc
	GetMessage() gin.HandlerFunc
	GetAllMessagesByUser() gin.HandlerFunc
	UpdateMessage() gin.HandlerFunc
	DeleteMessage() gin.HandlerFunc
}

type messageHandler struct {
	messageService service.MessageService
}

func NewMessageHandler(messageService service.MessageService) MessageHandler {
	return &messageHandler{
		messageService: messageService,
	}
}

func (s *messageHandler) CreateMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var message dto.MessageDto
		err := c.ShouldBindJSON(&message)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		err = s.messageService.Create(message)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		response := map[string]string{
			"status": "success",
			"data":   "new message created",
		}
		c.JSON(http.StatusOK, response)
	}
}

func (s *messageHandler) GetMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("handler error(incorrect message id): %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		message, err := s.messageService.Get(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		c.JSON(http.StatusOK, message)
	}
}

func (s *messageHandler) GetAllMessagesByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		messages, err := s.messageService.GetAllByUser(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		c.JSON(http.StatusOK, messages)
	}
}

func (s *messageHandler) UpdateMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var message dto.MessageDto
		err := c.ShouldBindJSON(&message)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		err = s.messageService.Update(message)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		response := map[string]string{
			"status": "success",
			"data":   "message updated",
		}
		c.JSON(http.StatusOK, response)
	}
}

func (s *messageHandler) DeleteMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("handler error(incorrect message id): %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		err = s.messageService.Delete(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		response := map[string]string{
			"status": "success",
			"data":   "message deleted",
		}
		c.JSON(http.StatusOK, response)
	}
}
