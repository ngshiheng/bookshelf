package app

import (
	"bookshelf/util/logger"

	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

type App struct {
	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate
}

func New(
	logger *logger.Logger,
	db *gorm.DB,
	validator *validator.Validate,
) *App {
	return &App{
		logger:    logger,
		db:        db,
		validator: validator,
	}
}

func (app *App) Logger() *logger.Logger {
	return app.logger
}

const (
	appErrDataAccessFailure   = "data access failure"
	appErrJsonCreationFailure = "json creation failure"
)

const (
	appErrDataCreationFailure = "data creation failure"
	appErrFormDecodingFailure = "form decoding failure"
)

const (
	appErrDataUpdateFailure = "data creation failure"
)
