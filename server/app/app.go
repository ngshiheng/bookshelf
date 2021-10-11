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
