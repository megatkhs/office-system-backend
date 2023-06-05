package bootstrap

import (
	"backend/utils/logger"

	"gorm.io/gorm"
)

type Application struct {
	Env *Env
	DB  *gorm.DB
}

func NewApp() *Application {
	logger.Initialize()

	app := Application{}

	app.Env = NewEnv()
	app.DB = NewDB(app.Env)

	return &app
}

func (app *Application) CloseDB() {
	CloseDB(app.DB)
}
