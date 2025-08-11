package db

import (
	"bitcoinmonitor/internal/db/model"
	"log/slog"
	"time"

	"github.com/harkaitz/go-coingecko"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func (s *Store) AddCurrencyToMonitoring(name string) error {
	tx := s.db.Update("monitoring", true).Where("name = ?", name)
	if tx.Error != nil {
		slog.Error("Error adding currency to monitoring", tx.Error.Error())
		return tx.Error
	}
	slog.Info("Currency added to monitoring", "name", name)
	return nil
}

func (s *Store) AddBitcoinPrice(name string, price float64) error {
	var currency model.AvailableCurrencyMonitoring
	tx := s.db.Where("name = ?", name).Take(&currency)
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
	slog.Info("Bitcoin price added to database", "name", name, "price", price)
	return nil
}

func (s *Store) GetMonitoringCoins() ([]model.AvailableCurrencyMonitoring, error) {
	var coins []model.AvailableCurrencyMonitoring
	if err := s.db.Select("name").Find(&coins).Where("monitoring = ?", true).Error; err != nil {
		slog.Error("Error getting monitoring coins from database", err.Error())
		return nil, err
	}
	slog.Info("Monitoring coins retrieved from database", "coins")
	return coins, nil
}

func (s *Store) GetAvailableCoins() ([]string, error) {
	var Coins []model.AvailableCurrencyMonitoring
	tx := s.db.Find(&Coins)
	if tx.Error != nil {
		slog.Error("Error getting available coins from database", tx.Error.Error())
		return nil, tx.Error
	}
	var names []string
	for _, currency := range Coins {
		names = append(names, currency.Name)
	}
	slog.Info("Available coins retrieved from database", "coins", names)
	return names, nil
}

func (s *Store) GetBitcoinPriceByName(name string, timestamp time.Time) (model.BitcoinPrice, error) {
	var price model.BitcoinPrice
	tx := s.db.
		Where("name = ?", name).
		Order("timestamp DESC").
		Where("timestamp <= ?", timestamp).
		Take(&price)
	if tx.Error != nil {
		slog.Error("Error getting bitcoin price from database", tx.Error.Error())
		return price, tx.Error
	}
	slog.Info("Bitcoin price retrieved from database", "name", name, "price", price.Price)
	return price, nil
}

func (s *Store) RemoveCurrencyFromMonitoring(name string) error {
	tx := s.db.Update("monitoring", false).Where("name = ?", name)
	if tx.Error != nil {
		slog.Error("Error removing currency from monitoring", tx.Error.Error())
		return tx.Error
	}
	slog.Info("Currency removed from monitoring", "name", name)
	return nil
}

func (s *Store) AddAvailableCoins(coins []coingecko.Coin) error {
	for _, coin := range coins {
		s.db.Create(&model.AvailableCurrencyMonitoring{CoinID: string(coin.ID), Name: coin.Symbol, Description: coin.Name})
	}
	slog.Info("Available coins added to database", "coins")
	return nil
}
