package routes

import (
	"backend/bootstrap"
	"backend/utils/logger"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

func Setup(app *bootstrap.App, fiberApp *fiber.App) {
	api := fiberApp.Group("/api")
	v1 := api.Group("/v1")

	api.Use(newLogger)

	setupEmployeeRoutes(app.DB, v1)
	setupDepartmentRoutes(app.DB, v1)
}

func newLogger() fiber.Handler {
	tz, _ := time.LoadLocation("Asia/Tokyo")

	return func(c *fiber.Ctx) error {
		start := time.Now().In(tz)

		// リクエスト
		err := c.Next()

		msg := "Request"
		if err != nil {
			msg = err.Error()

			switch {
			case errors.Is(err, fiber.ErrNotFound):
				c.SendStatus(fiber.ErrNotFound.Code)
			case errors.Is(err, fiber.ErrInternalServerError):
				c.SendStatus(fiber.ErrInternalServerError.Code)
			}
		}

		code := c.Response().StatusCode()
		end := time.Now().In(tz)

		attrs := []slog.Attr{
			slog.String("path", c.Path()),
			slog.String("method", c.Method()),
			slog.Int("status", code),
			slog.Int64("latency", end.Sub(start).Milliseconds()),
			slog.String("ip", c.IP()),
			slog.String("user_agent", string(c.Request().Header.UserAgent())),
		}

		logger.Info(msg, attrs...)

		return err
	}
}
