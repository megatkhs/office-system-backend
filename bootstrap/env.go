package bootstrap

import (
	"backend/utils/logger"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	AppEnv     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBHost     string
	APIPort    string
}

func NewEnv() *Env {
	env := Env{}

	if err := godotenv.Load(); err != nil {
		logger.Fatal(".envを読み込めませんでした", err)
	}

	env.AppEnv = os.Getenv("APP_ENV")
	env.DBUser = os.Getenv("DB_USER")
	env.DBPassword = os.Getenv("DB_PASSWORD")
	env.DBName = os.Getenv("DB_NAME")
	env.DBPort = os.Getenv("DB_PORT")
	env.DBHost = os.Getenv("DB_HOST")
	env.APIPort = os.Getenv("API_PORT")

	logger.Info(fmt.Sprintf("%sモードで動作中", env.AppEnv))

	return &env
}
