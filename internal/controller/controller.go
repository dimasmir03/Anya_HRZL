package controller

import (
	"net/http"

	"bitcoinmonitor/internal/service"

	"github.com/labstack/echo/v4"
)

type BitcoinPriceRequest struct {
	Name      string `json:"name" validate:"required"`
	Timestamp int64  `json:"timestamp" validate:"required"`
}

type BitcoinController struct {
	service *service.BitcoinService
}

func NewBitcoinController(service *service.BitcoinService) *BitcoinController {
	return &BitcoinController{
		service: service,
	}
}

func (c *BitcoinController) GetBitcoinPrice(e echo.Context) error {
	var req BitcoinPriceRequest
	if err := e.Bind(&req); err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}
	price, err := c.service.GetBitcoinPriceByName(req.Name, req.Timestamp)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, price)
}

func (c *BitcoinController) AddCurrencyToMonitoring(e echo.Context) error {
	var req BitcoinPriceRequest
	if err := e.Bind(&req); err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}

	if err := c.service.AddCurrencyToMonitoring(req.Name); err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.String(http.StatusOK, "currency added to monitoring")
}

func (c *BitcoinController) RemoveCurrencyFromMonitoring(e echo.Context) error {
	var req BitcoinPriceRequest
	if err := e.Bind(&req); err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}

	if err := c.service.RemoveCurrencyFromMonitoring(req.Name); err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.String(http.StatusOK, "currency removed from monitoring")
}

func (c *BitcoinController) GetMonitoringCoins(e echo.Context) error {
	Coins, err := c.service.GetMonitoringCoins()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, Coins)
}

func (c *BitcoinController) GetAvailableCoins(e echo.Context) error {
	Coins, err := c.service.GetAvailableCoins()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, Coins)
}

func (c *BitcoinController) StartTimer(e echo.Context) error {
	if err := c.service.StartTimer(); err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.String(http.StatusOK, "monitoring started")
}

func (c *BitcoinController) StopTimer(e echo.Context) error {
	if err := c.service.StopTimer(); err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.String(http.StatusOK, "monitoring stopped")
}
