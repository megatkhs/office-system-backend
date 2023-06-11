package main

import (
	"backend/bootstrap"
	"backend/model"
	"backend/utils/logger"
)

func main() {
	app := bootstrap.NewApp()

	app.DB.AutoMigrate(model.Employee{}, model.Contact{}, model.Department{})
	logger.Info("マイグレーションが成功しました")

	app.CloseDB()
}
