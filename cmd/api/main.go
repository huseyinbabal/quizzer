package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huseyinbabal/quizzer/internal/config"
	"github.com/huseyinbabal/quizzer/internal/domain/question"
	"github.com/huseyinbabal/quizzer/internal/loggerx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {

	logger, err := loggerx.New()
	if err != nil {
		log.Fatalf("failed to initialize logger %v", err)
	}

	cfg, err := config.Init()
	if err != nil {
		logger.Fatal("failed to initialize config", zap.Error(err))
	}

	db, err := gorm.Open(postgres.Open(cfg.DB.Dsn()), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database", zap.Error(err))

	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("failed to create database instance", zap.Error(err))

	}

	defer sqlDB.Close()

	app := fiber.New()

	q, err := question.New(db, logger)
	if err != nil {
		logger.Fatal("failed to initialize question repository", zap.Error(err))
	}
	app.Get("/questions", func(ctx *fiber.Ctx) error {
		questions := q.List(ctx.Context())
		err := ctx.JSON(questions)
		return err
	})

	err = app.Listen(":3000")
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}
}
