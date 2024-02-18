package router

import (
	"encoding/json"
	"time"

	"github.com/giovanniussuy/gdev-go-lang/infra/config"
	"github.com/gofiber/fiber/v2"
)

type (
	Server interface {
		Listen()
	}
)

type (
	ServerEngine struct {
		appName        string
		logLevel       string
		middleware     *fiber.App
		port           int64
		contextTimeout time.Duration
	}
)

func StartWebEngine() *ServerEngine {
	return &ServerEngine{
		middleware: fiber.New(fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
	}
}

func (engine *ServerEngine) WithAppName() *ServerEngine {
	engine.appName = config.C.AppName
	engine.logLevel = config.C.LoggingLevel
}
