package application

import (
	"bitcoinmonitor/internal/config"
	"bitcoinmonitor/internal/controller"
	"bitcoinmonitor/internal/db"
	"bitcoinmonitor/internal/service"
	"bitcoinmonitor/internal/transport"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Application struct {
	e *echo.Echo
}

func NewApplication() *Application {
	return &Application{e: echo.New()}
}

func (a *Application) Run() {
	config_path := ".env"
	cfg := config.LoadСonfig(config_path)
	store, err := db.NewStore(cfg.DB_Url)
	if err != nil {
		panic(err)
	}
	service := service.NewBitcoinService(store)
	service.StartMonitoring()
	ctrl := controller.NewBitcoinController(service)
	a.e.Use(middleware.Logger())
	transport.NewRouter(a.e, ctrl)
	a.e.Logger.Fatal(a.e.Start(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)))

}
