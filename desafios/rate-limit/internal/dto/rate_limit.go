package dto

import (
	"github.com/felipedsf/fullcycle-trannings/desafios/rate-limit/internal/config"
	"strconv"
	"strings"
	"time"
)

type RateLimit struct {
	Id       string
	Until    time.Time
	Attempts int

	MaxAttempts int
	Interval    int

	BlockedUntil  *time.Time
	BlockInterval int
}

func NewRateLimit(token string, ip string) *RateLimit {

	var id string
	cfg := config.AppConfig

	configuredInterval := cfg.Interval
	configuredBlockInterval := cfg.BlockInterval
	configuredLimit := cfg.Limit

	if tc, ok := cfg.TokenConfig[token]; ok {
		id = token
		values := strings.Split(tc, "_")
		if len(values) >= 3 {
			if parsedInterval, err := parseTokenConfigValue(values[0]); err == nil {
				configuredInterval = parsedInterval
			}
			if parsedBlockInterval, err := parseTokenConfigValue(values[1]); err == nil {
				configuredBlockInterval = parsedBlockInterval
			}
			if parsedLimit, err := parseTokenConfigValue(values[2]); err == nil {
				configuredLimit = parsedLimit
			}
		}
	} else {
		id = ip
	}

	return &RateLimit{
		Id:            id,
		Attempts:      1,
		Until:         time.Now().Add(time.Duration(configuredInterval) * time.Second),
		Interval:      configuredInterval,
		BlockInterval: configuredBlockInterval,
		MaxAttempts:   configuredLimit,
	}
}

func parseTokenConfigValue(value string) (int, error) {
	return strconv.Atoi(value)
}
