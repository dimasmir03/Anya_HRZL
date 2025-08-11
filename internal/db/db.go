package db

import (
	"bitcoinmonitor/internal/db/model"
	"fmt"
	"log/slog"
	"time"

	"github.com/harkaitz/go-coingecko"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Store struct {
	db *gorm.DB
}

func NewStore(url string) (*Store, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		slog.Error("Error connecting to database", err.Error())
		return nil, err
	}

	if err := initModels(db); err != nil {
		slog.Error("Error initializing database models", err.Error())
		return nil, err
	}

	return &Store{db: db}, nil
}

func initModels(db *gorm.DB) error {
	models := []interface{}{
		&model.AvailableCurrencyMonitoring{},
		&model.BitcoinPrice{},
	}
	if err := db.AutoMigrate(models...); err != nil {
		slog.Error("Error auto-migrating database models", err.Error())
		return err
	}

	return nil
}

func (s *Store) AddCoinToMonitoring(name string) error {
	tx := s.db.Table("available_currency_monitorings").Where("name = ?", name).Update("monitoring", true)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("Currency not found in database", "name", name)
		}
		return fmt.Errorf("Error adding currency to monitoring", tx.Error.Error())
	}
	return nil
}

func (s *Store) AddBitcoinPrice(id string, price float64) error {
	var currency model.AvailableCurrencyMonitoring
	tx := s.db.Table("available_currency_monitorings").Where("coin_id = ?", id).Take(&currency)
	if tx.Error != nil {
		slog.Error("Error getting currency from database", tx.Error.Error())
		return tx.Error
	}
	priceModel := model.BitcoinPrice{BitcoinID: currency.ID, Price: price}
	tx = s.db.Create(&priceModel)
	if tx.Error != nil {
		slog.Error("Error adding bitcoin price to database", tx.Error.Error())
		return tx.Error
	}
	slog.Info("Bitcoin price added to database", "coin_id", id, "price", price)
	return nil
}

func (s *Store) GetMonitoringCoins() ([]model.AvailableCurrencyMonitoring, error) {
	var coins []model.AvailableCurrencyMonitoring
	if err := s.db.Select("*").Where("monitoring = ?", true).Find(&coins).Error; err != nil {
		slog.Error("Error getting monitoring coins from database", err.Error())
		return nil, err
	}
	slog.Info("Monitoring coins retrieved from database", "coins")
	return coins, nil
}

func (s *Store) GetAvailableCoins() ([]model.AvailableCurrencyMonitoring, error) {
	var Coins []model.AvailableCurrencyMonitoring
	tx := s.db.Find(&Coins)
	if tx.Error != nil {
		slog.Error("Error getting available coins from database", tx.Error.Error())
		return nil, tx.Error
	}
	slog.Info("Available coins retrieved from database", "coins")
	return Coins, nil
}

func (s *Store) GetBitcoinPriceByName(name string, timestamp time.Time) (model.BitcoinPrice, error) {
	// var price model.BitcoinPrice
	// tx := s.db.
	// 	Where("name = ?", name).
	// 	Order("timestamp DESC").
	// 	Where("timestamp <= ?", timestamp).
	// 	Take(&price)
	// if tx.Error != nil {
	// 	slog.Error("Error getting bitcoin price from database", tx.Error.Error())
	// 	return price, tx.Error
	// }
	// slog.Info("Bitcoin price retrieved from database", "name", name, "price", price.Price)
	// return price, nil

	var price model.BitcoinPrice
	tx := s.db.Table("bitcoin_prices").
		Clauses(clause.OrderBy{
			Expression: clause.Expr{
				SQL:  "ABS(TIMESTAMPDIFF(SECOND, created_at, ?))",
				Vars: []interface{}{timestamp},
			},
		}).
		First(&price)
	if tx.Error != nil {
		slog.Error("Error getting bitcoin price from database", tx.Error.Error())
		return price, tx.Error
	}
	slog.Info("Bitcoin price retrieved from database", "name", name, "price", price.Price)
	return price, nil
}

func (s *Store) RemoveCoinFromMonitoring(name string) error {
	tx := s.db.Table("available_currency_monitorings").Where("name = ?", name).Update("monitoring", false)
	if tx.Error != nil {
		slog.Error("Error removing currency from monitoring", tx.Error.Error())
		return tx.Error
	}
	slog.Info("Currency removed from monitoring", "name", name)
	return nil
}

func (s *Store) AddAvailableCoins(coins []coingecko.Coin) error {
	for _, coin := range coins {
		var existing model.AvailableCurrencyMonitoring
		tx := s.db.Where("coin_id = ?", coin.ID).Take(&existing)
		if tx.Error == gorm.ErrRecordNotFound {
			s.db.Create(&model.AvailableCurrencyMonitoring{CoinID: string(coin.ID), Symbol: coin.Symbol, Name: coin.Name})
		}
	}
	slog.Info("Available coins added to database", "coins")
	return nil
}
