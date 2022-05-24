package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"messenger/pkg/handler"
)

type Server struct {
	router         *gin.Engine
	userHandler    handler.UserHandler
	messageHandler handler.MessageHandler
}

func NewServer(router *gin.Engine, userHandler handler.UserHandler, messageHandler handler.MessageHandler) *Server {
	return &Server{
		router:         router,
		userHandler:    userHandler,
		messageHandler: messageHandler,
	}
}

func (s *Server) Run() error {
	r := s.Routes()
	err := r.Run()
	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}
	return nil
}
