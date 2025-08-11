package transport

import (
	"bitcoinmonitor/internal/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, ctrl *controller.BitcoinController) *echo.Echo {

	e.POST("/bitcoin/add", ctrl.AddCurrencyToMonitoring)
	e.POST("/bitcoin/remove", ctrl.RemoveCurrencyFromMonitoring)
	e.GET("/bitcoin/price", ctrl.GetBitcoinPrice)
	e.GET("/bitcoin/monitoring", ctrl.GetMonitoringCoins)
	e.GET("/bitcoin/available", ctrl.GetAvailableCoins)
	e.GET("/startmonitoring", ctrl.StartTimer)
	e.GET("/stopmonitoring", ctrl.StopTimer)

	return e
}
