package main

import (
	"backend/bootstrap"
	"backend/routes"
	"backend/utils/logger"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := bootstrap.NewApp()
	defer app.CloseDB()

	f := fiber.New()

	f.Use(cors.New())

	routes.Setup(app, f)

	p := fmt.Sprintf(":%s", app.Env.APIPort)
	logger.Fatal("APIサーバーを起動しました", f.Listen(p))
}
