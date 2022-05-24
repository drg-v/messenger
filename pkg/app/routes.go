package app

import "github.com/gin-gonic/gin"

func (s *Server) Routes() *gin.Engine {
	router := s.router
	users := router.Group("/users")
	{
		users.GET("/:id", s.userHandler.GetUser())
		users.GET("", s.userHandler.GetAllUsers())
		users.POST("", s.userHandler.CreateUser())
		users.PUT("", s.userHandler.UpdateUser())
		users.DELETE("/:id", s.userHandler.DeleteUser())
	}
	messages := router.Group("/messages")
	{
		messages.GET("/:id", s.messageHandler.GetAllMessagesByUser())
		messages.POST("", s.messageHandler.CreateMessage())
		messages.PUT("", s.messageHandler.UpdateMessage())
		messages.DELETE("/:id", s.messageHandler.DeleteMessage())
	}
	return router
}
