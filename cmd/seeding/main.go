package main

import (
	"backend/bootstrap"
	"backend/utils/logger"

	"gorm.io/gorm"
)

func main() {
	app := bootstrap.NewApp()

	if err := seeds(app.DB); err != nil {
		app.CloseDB()
		logger.Fatal("seedingに失敗しました", err)
	}

	logger.Info("seedingに成功しました")
	app.CloseDB()
}

func seeds(db *gorm.DB) error {
	return nil
}
