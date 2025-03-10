package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	cfgPath string = "D:/thebestprojects/URLShortener/config/local.yaml"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}
type HTTPServer struct {
	Addres      string        `yaml:"addres" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func InitConfig() *Config {
	configPath := cfgPath
	if configPath == "" {
		log.Fatal("Config path not found")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file:%s exists", configPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("File:%s is not read\n, Error:%s", configPath, err.Error())
	}
	return &cfg
}
