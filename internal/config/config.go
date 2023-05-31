package config

import (
	"sync"

	"github.com/alexPavlikov/REST-API_Clean_Architecture/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

var (
	Path = "./config.cfg"
)

type Config struct {
	IsDebug bool `json:"is_debug"`
	Listen  struct {
		Type   string `json:"type"`
		BindIP string `json:"bind_ip"`
		Port   string `json:"port"`
	} `json:"listen"`
	Storage StorageConfig `json:"storage"`
}

type StorageConfig struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	Database    string `json:"database"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	MaxAttempts string `json:"max_attempts"`
}

var cfg *Config
var once sync.Once

func GetConfig() *Config {
	logger := logging.GetLogger()
	once.Do(func() {
		cfg = &Config{}
		err := cleanenv.ReadConfig("./internal/config/config.yml", cfg)
		if err != nil {
			help, _ := cleanenv.GetDescription(cfg, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return cfg
}
