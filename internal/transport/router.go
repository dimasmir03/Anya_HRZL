package transport

import (
	"bitcoinmonitor/internal/controller"

	"github.com/labstack/echo/v4"

	_ "bitcoinmonitor/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(e *echo.Echo, ctrl *controller.BitcoinController) *echo.Echo {

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	g := e.Group("/api/v1")
	g.POST("/bitcoin/add", ctrl.AddCoinToMonitoring)
	g.POST("/bitcoin/remove", ctrl.RemoveCurrencyFromMonitoring)
	g.GET("/bitcoin/price", ctrl.GetBitcoinPrice)
	g.GET("/bitcoin/monitoring", ctrl.GetMonitoringCoins)
	g.GET("/bitcoin/available", ctrl.GetAvailableCoins)
	g.GET("/startmonitor", ctrl.StartTimer)
	g.GET("/stopmonitor", ctrl.StopTimer)

	return e
}
