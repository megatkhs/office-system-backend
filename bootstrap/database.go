package bootstrap

import (
	"backend/utils/logger"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(env *Env) *gorm.DB {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", env.DBUser, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		logger.Fatal("データベースとの接続を確立できませんでした", err)
	}

	logger.Info("データベースとの接続を確立")

	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()

	if err := sqlDB.Close(); err != nil {
		logger.Fatal("データベースとの接続を切断できませんでした", err)
	}

	logger.Info("データベースとの接続を切断")
}
