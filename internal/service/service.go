package service

import (
	"bitcoinmonitor/internal/db"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/harkaitz/go-coingecko"
)

type BitcoinService struct {
	store          *db.Store
	ticker         *time.Ticker
	tickerDuration time.Duration
	monitoringList []string
}

func NewBitcoinService(store *db.Store) *BitcoinService {
	tickerDuration := 15 * time.Second
	return &BitcoinService{
		store:          store,
		tickerDuration: tickerDuration,
		ticker:         time.NewTicker(tickerDuration),
	}
}

func (s *BitcoinService) UpdateMonitoringList() error {
	monitoring_coins, err := s.store.GetMonitoringCoins()
	if err != nil {
		return err
	}
	s.monitoringList = make([]string, len(monitoring_coins))
	for i, coin := range monitoring_coins {
		s.monitoringList[i] = coin.Name
	}
	slog.Info("Monitoring list updated", "list", s.monitoringList)
	return nil

}

type BitcoinPrice struct {
	Price float64 `json:"usd"`
}

func (s *BitcoinService) InitMonitoringData() error {
	slog.Info("Bitcoin Service starting")
	defer s.ticker.Stop()

	coins, err := coingecko.GetCoinList()
	if err != nil {
		return err
	}
	if err := s.store.AddAvailableCoins(coins); err != nil {
		return err
	}

	s.UpdateMonitoringList()

	go func() {
		slog.Info("Monitoring started")
		for range s.ticker.C {
			if len(s.monitoringList) == 0 {
				continue
			}
			slog.Info("Monitoring tick")
			resp, err := http.Get(fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", strings.Join(s.monitoringList, ",")))
			if err != nil {
				slog.Error("Error getting bitcoin price", err.Error())
				continue
			}

			if resp.StatusCode != http.StatusOK {
				slog.Error("unexpected status code: %d", strconv.Itoa(resp.StatusCode))
				continue
			}
			var data map[string]BitcoinPrice
			if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
				slog.Error("Error decoding bitcoin price", err.Error())
				continue
			}
			log.Println(data)
			resp.Body.Close()
			for id, price := range data {
				if err := s.store.AddBitcoinPrice(id, price.Price); err != nil {
					slog.Error("Error adding bitcoin price", err.Error())
					continue
				}
			}
		}
	}()

	return nil
}

func (s *BitcoinService) GetBitcoinPriceByName(name string, timestamp int64) (float64, error) {
	d, e := s.store.GetBitcoinPriceByName(name, time.Unix(timestamp, 0))
	return d.Price, e
}

func (s *BitcoinService) AddCurrencyToMonitoring(name string) error {
	if err := s.store.AddCoinToMonitoring(name); err != nil {
		return err
	}
	log.Printf("coin %s added to monitoring\n", name)

	if err := s.UpdateMonitoringList(); err != nil {
		return err
	}
	return nil
}

func (s *BitcoinService) RemoveCurrencyFromMonitoring(name string) error {
	if err := s.store.RemoveCoinFromMonitoring(name); err != nil {
		return err
	}
	if err := s.UpdateMonitoringList(); err != nil {
		return err
	}
	return nil
}

func (s *BitcoinService) GetMonitoringCoins() ([]string, error) {
	coins, err := s.store.GetMonitoringCoins()
	if err != nil {
		return nil, err
	}
	var names []string
	for _, coin := range coins {
		names = append(names, coin.Name)
	}
	return names, nil
	// return coins, nil
}

func (s *BitcoinService) GetAvailableCoins() ([]string, error) {
	c, err := s.store.GetAvailableCoins()
	if err != nil {
		return nil, err
	}
	var names []string
	for _, coin := range c {
		names = append(names, coin.Name)
	}
	return names, nil
}

func (s *BitcoinService) StartTimer() error {
	s.ticker.Reset(s.tickerDuration)
	return nil
}

func (s *BitcoinService) StopTimer() error {
	s.ticker.Stop()
	return nil
}
