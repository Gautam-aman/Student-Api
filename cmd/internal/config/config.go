package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string `yaml:"address"`
}

type Config struct {
	Env          string `yaml:"env"`
	Storage_path string `yaml:"storage_path" env-required:"true"`
	HttpServer   `yaml:"http_server"`
}

func MustLoad() *Config {

	var ConfigPath string

	ConfigPath = os.Getenv("CONFIG_PATH")

	if ConfigPath == "" {
		flags := flag.String("config", "", "path to config is missing")
		flag.Parse()
		ConfigPath = *flags
	}
	if ConfigPath == "" {
		log.Fatal("Missing config path")
	}

	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		log.Fatalf("Config does not exist: %s", ConfigPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(ConfigPath, &cfg)

	if err != nil {
		log.Fatalf("Cannot read config file: %s", err.Error())
	}

	return &cfg

}
