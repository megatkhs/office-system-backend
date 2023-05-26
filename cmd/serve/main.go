package main

import (
	"backend/bootstrap"
	"backend/utils/logger"
)

func main() {
	app := bootstrap.NewApp()
	defer app.CloseDB()

	logger.Info("APIサーバーを起動しました")
}
