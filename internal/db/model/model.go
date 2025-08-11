package model

import "time"

type AvailableCurrencyMonitoring struct {
	ID           int64          `json:"id" gorm:"primary_key;tableName:available_currency_monitoring"`
	CoinID       string         `json:"coin_id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Monitoring   bool           `json:"monitoring" gorm:"default:false"`
	BitcoinPrice []BitcoinPrice `json:"bitcoin_prices_id" gorm:"foreignKey:bitcoin_id"`
}

type BitcoinPrice struct {
	ID        int64     `json:"id" gorm:"primary_key;tableName:bitcoin_price"`
	BitcoinID int64     `json:"bitcoin_id"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP"`
}
