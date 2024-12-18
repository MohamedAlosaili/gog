package middleware

import (
	"fmt"

	"github.com/PROJECT_NAME/internal/logger"
	"github.com/gofiber/fiber/v2"
)

type (
	loggingMiddlewareDependencies interface {
		logger.LoggerProvider
	}

	LoggingMiddleware struct {
		d loggingMiddlewareDependencies
	}
)

func NewLoggingMiddleware(d loggingMiddlewareDependencies) *LoggingMiddleware {
	return &LoggingMiddleware{d: d}
}

func (m *LoggingMiddleware) Handle(c *fiber.Ctx) error {
	m.d.Logger().Info(fmt.Sprintf("A new request path: %s query: %s method: %s request_id: %s ip: %s user_agent: %s content_length: %d",
		c.Path(),
		string(c.Request().URI().QueryString()),
		c.Method(),
		c.Locals("RequestID"),
		c.IP(),
		string(c.Request().Header.UserAgent()),
		int64(len(c.Body())),
	))

	return c.Next()
}
