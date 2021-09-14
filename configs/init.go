package configs

import (
	_ "embed"
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/sharpvik/log-go/v2"
)

//go:embed default.yml
var defaultConfig []byte

// These values are exposed to every other package for reading.
var (
	Database *DatabaseConfig
	Server   *ServerConfig
)

func init() {
	log.Debug("reading config ...")

	config := new(Config)

	if isTest := os.Getenv("TEST"); len(isTest) > 0 {
		readDefaultConfig(config)
		setValues(config)
		return
	}

	flags := parseFlags()
	if *flags.ConfigPath == "" {
		readDefaultConfig(config)
	} else {
		readCustomConfig(*flags.ConfigPath, config)
	}

	setValues(config)

	log.Debug("config successfull")
}

func readDefaultConfig(config *Config) {
	if err := yaml.Unmarshal(defaultConfig, config); err != nil {
		log.Fatal("failed to read default config file: %s", err)
	}
}

func readCustomConfig(name string, config *Config) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal("failed to open custom config file")
	}
	if err := json.NewDecoder(file).Decode(config); err != nil {
		log.Fatal("failed to read default config file: %s", err)
	}
}

func setValues(config *Config) {
	Database = config.Database
	Server = config.Server
}
