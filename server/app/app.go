package app

import (
	"bookshelf/util/logger"

	"github.com/jinzhu/gorm"
)

type App struct {
	logger *logger.Logger
	db     *gorm.DB
}

func New(
	logger *logger.Logger,
	db *gorm.DB,
) *App {
	return &App{
		logger: logger,
		db:     db,
	}
}
