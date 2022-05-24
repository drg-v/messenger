package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"messenger/pkg/dto"
	"messenger/pkg/service"
	"net/http"
	"strconv"
)

type UserHandler interface {
	CreateUser() gin.HandlerFunc
	GetUser() gin.HandlerFunc
	GetAllUsers() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (s *userHandler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var user dto.UserDto
		err := c.ShouldBindJSON(&user)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		err = s.userService.Create(user)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		response := map[string]string{
			"status": "success",
			"data":   "new user created",
		}
		c.JSON(http.StatusOK, response)
	}
}

func (s *userHandler) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("handler error(incorrect user id): %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		user, err := s.userService.Get(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func (s *userHandler) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		users, err := s.userService.GetAll()
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func (s *userHandler) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var user dto.UserDto
		err := c.ShouldBindJSON(&user)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		err = s.userService.Update(user)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		response := map[string]string{
			"status": "success",
			"data":   "user updated",
		}
		c.JSON(http.StatusOK, response)
	}
}

func (s *userHandler) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("handler error(incorrect user id): %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		err = s.userService.Delete(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		response := map[string]string{
			"status": "success",
			"data":   "user deleted",
		}
		c.JSON(http.StatusOK, response)
	}
}
