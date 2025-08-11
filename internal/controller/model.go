package controller

type BitcoinPriceRequest struct {
	CointName string `json:"coin" validate:"required"`
	Timestamp int64  `json:"timestamp" validate:"required"`
}

type BitcoinPriceResponse struct {
	Price float64 `json:"price"`
}

type BitcoinAddRequest struct {
	CoinName string `json:"coin" validate:"required"`
}

type BitcoinRemoveRequest struct {
	CoinName string `json:"coin" validate:"required"`
}

type MonitoringCoinsResponse struct {
	Coins []string `json:"coins"`
}

type AvailableCoinsResponse struct {
	Coins []string `json:"coins"`
}
