package bootstrap

import (
	"backend/utils/logger"

	"gorm.io/gorm"
)

type App struct {
	Env *Env
	DB  *gorm.DB
}

func NewApp() *App {
	logger.Initialize()

	app := App{}

	app.Env = NewEnv()
	app.DB = NewDB(app.Env)

	return &app
}

func (app *App) CloseDB() {
	CloseDB(app.DB)
}
