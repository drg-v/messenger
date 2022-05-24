package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"messenger/pkg/app"
	"messenger/pkg/entity"
	"messenger/pkg/handler"
	"messenger/pkg/repository"
	"messenger/pkg/service"
	"messenger/util"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\\n", err)
		os.Exit(1)
	}
}

func run() error {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432",
		config.ServerAddress,
		config.DBUsername,
		config.DBPassword,
		config.DBName,
	)
	db, err := setupDatabase(dsn)
	if err != nil {
		return err
	}
	userRepository := repository.NewUserRepository(db)
	messageRepository := repository.NewMessageRepository(db)
	router := gin.Default()
	router.Use(cors.Default())
	userService := service.NewUserService(userRepository)
	messageService := service.NewMessageService(messageRepository)
	userHandler := handler.NewUserHandler(userService)
	messageHandler := handler.NewMessageHandler(messageService)
	server := app.NewServer(router, userHandler, messageHandler)
	err = server.Run()
	if err != nil {
		return err
	}
	return nil
}

func setupDatabase(connection string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&entity.User{}, &entity.Message{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
