package router

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/giovanniussuy/gdev-go-lang/app/controller"
	"github.com/giovanniussuy/gdev-go-lang/infra/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	return engine
}

func (engine *ServerEngine) WithPort(port int64) *ServerEngine {
	engine.port = port
	return engine
}

func (engine *ServerEngine) WithDuration(timeout time.Duration) *ServerEngine {
	engine.contextTimeout = timeout
	return engine
}

func (engine *ServerEngine) Listen() {
	engine.middleware.Use(recover.New())
	engine.WithRoutes(engine.middleware)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", engine.port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler:      adaptor.FiberApp(engine.middleware),
	}

	notifyContext, cancelFunc := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cancelFunc()

	waitGroup := &sync.WaitGroup{}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Error to start and serve the Web Server Engine!", err)
		}
	}()

	<-notifyContext.Done()

	cancelContext, cancelFunc := context.WithTimeout(notifyContext, 10*time.Second)
	defer cancelFunc()

	waitGroup.Wait()

	if err := server.Shutdown(cancelContext); err != nil {
		log.Fatal("Error to shutdown Server Engine!", err)
	}
}

func (engine *ServerEngine) WithRoutes(routeEngine *fiber.App) {
	routeEngine.Post("/v1/nothing", controller.NothingController)
}
