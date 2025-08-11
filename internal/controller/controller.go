package controller

import (
	"fmt"
	"log"
	"net/http"

	"bitcoinmonitor/internal/service"

	"github.com/labstack/echo/v4"
)

type BitcoinController struct {
	service *service.BitcoinService
}

func NewBitcoinController(service *service.BitcoinService) *BitcoinController {
	return &BitcoinController{
		service: service,
	}
}

// GetBitcoinPrice - получает курс биткойна по имени валюты и timestamp
// @Summary Получает курс биткойна по имени валюты и timestamp
// @Produce json
// @Param coin path string true "Имя монеты"
// @Param timestamp path integer true "момент времени"
// @Success 200 {object} float64
// @Router /bitcoin/price [get]
func (c *BitcoinController) GetBitcoinPrice(e echo.Context) error {
	var req BitcoinPriceRequest
	if err := e.Bind(&req); err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}
	price, err := c.service.GetBitcoinPriceByName(req.CointName, req.Timestamp)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, &BitcoinPriceResponse{
		Price: price,
	})
}

// AddCurrencyToMonitoring - добавляет валюту в список мониторинга
// @Summary Добавляет валюту в список мониторинга
// @Produce json
// @Accept json
// @Param coin body string true "Имя монеты"
// @Success 200 {string} string
// @Router /bitcoin/add [post]
func (c *BitcoinController) AddCoinToMonitoring(e echo.Context) error {
	var req BitcoinAddRequest
	if err := e.Bind(&req); err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}
	log.Println(req)
	if err := c.service.AddCurrencyToMonitoring(req.CoinName); err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.String(http.StatusOK, fmt.Sprintf("валюта %s добавлена в список мониторинга", req.CoinName))
}

// RemoveCurrencyFromMonitoring - удаляет валюту из списка мониторинга
// @Summary Удаляет валюту из списка мониторинга
// @Produce json
// @Param coin path string true "Имя монеты"
// @Success 200 {string} string
// @Router /bitcoin/remove [post]
func (c *BitcoinController) RemoveCurrencyFromMonitoring(e echo.Context) error {
	var req BitcoinRemoveRequest
	if err := e.Bind(&req); err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}

	if err := c.service.RemoveCurrencyFromMonitoring(req.CoinName); err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.String(http.StatusOK, fmt.Sprintf("валюта %s удалена из списка мониторинга", req.CoinName))
}

// GetMonitoringCoins - получает список мониторингуемых валют
// @Summary Получает список мониторингуемых валют
// @Produce json
// @Success 200 {array} model.AvailableCurrencyMonitoring
// @Router /bitcoin/monitoring [get]
func (c *BitcoinController) GetMonitoringCoins(e echo.Context) error {
	Coins, err := c.service.GetMonitoringCoins()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	log.Println(Coins)
	return e.JSON(http.StatusOK, Coins)
}

// GetAvailableCoins - получает список доступных валют
// @Summary Получает список доступных валют
// @Produce json
// @Success 200 {array} model.AvailableCurrencyMonitoring
// @Router /bitcoin/available [get]
func (c *BitcoinController) GetAvailableCoins(e echo.Context) error {
	Coins, err := c.service.GetAvailableCoins()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, Coins)
}

// StartTimer - стартует мониторинг
// @Summary Стартует мониторинг
// @Produce json
// @Success 200 {string} string
// @Router /startmonitor [get]
func (c *BitcoinController) StartTimer(e echo.Context) error {
	if err := c.service.StartTimer(); err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.String(http.StatusOK, "мониторинг запущен")
}

// StopTimer - останавливает мониторинг
// @Summary Останавливает мониторинг
// @Produce json
// @Success 200 {string} string
// @Router /stopmonitor [get]
func (c *BitcoinController) StopTimer(e echo.Context) error {
	if err := c.service.StopTimer(); err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.String(http.StatusOK, "мониторинг остановлен")
}
