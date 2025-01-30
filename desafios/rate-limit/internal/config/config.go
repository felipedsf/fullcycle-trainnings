package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var AppConfig *config

func init() {
	err := loadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
}

type config struct {
	Port          string            `mapstructure:"PORT"`
	Redis         string            `mapstructure:"REDIS"`
	TokenConfig   map[string]string `mapstructure:"-"`
	Limit         int               `mapstructure:"LIMIT"`
	Interval      int               `mapstructure:"INTERVAL"`
	BlockInterval int               `mapstructure:"BLOCK_INTERVAL"`
}

type rateConfig struct {
	Token         string `json:"TOKEN"`
	Limit         int    `json:"LIMIT"`
	Interval      int    `json:"INTERVAL"`
	BlockInterval int    `json:"BLOCK_INTERVAL"`
}

func loadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	var cfg config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	if cfg.Port == "" {
		log.Printf("default port is %s\n", "8080")
		cfg.Port = ":8080"
	}

	if cfg.BlockInterval == 0 || cfg.Limit == 0 || cfg.Interval == 0 {
		return errors.New("default limit, interval or block interval is not configured")
	}

	if viper.IsSet("TOKEN_CONFIG") {
		var rc []rateConfig
		rawTokenCfg := viper.GetString("TOKEN_CONFIG")
		err = json.Unmarshal([]byte(rawTokenCfg), &rc)
		if err != nil {
			return err
		}
		for _, rateConfig := range rc {
			if cfg.TokenConfig == nil {
				cfg.TokenConfig = make(map[string]string)
			}
			cfg.TokenConfig[rateConfig.Token] = fmt.Sprintf("%d_%d_%d", rateConfig.Interval, rateConfig.BlockInterval, rateConfig.Limit)
		}

	}
	AppConfig = &cfg
	return nil
}
