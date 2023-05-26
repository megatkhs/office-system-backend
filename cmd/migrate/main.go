package main

import (
	"backend/bootstrap"
	"backend/utils/logger"
)

func main() {
	app := bootstrap.NewApp()

	app.DB.AutoMigrate()
	logger.Info("マイグレーションが成功しました")

	app.CloseDB()
}
