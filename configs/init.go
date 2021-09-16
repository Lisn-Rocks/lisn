package configs

import (
	_ "embed"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/sharpvik/log-go/v2"
)

var (
	/*
	 * These values are exposed to every other package for reading.
	 ! But they must not be modified.
	*/
	Database *DatabaseConfig
	Server   *ServerConfig

	//go:embed default.yml
	defaultConfig []byte
	config        Config
)

func Init() {
	log.Debug("reading config ...")

	flags := parseFlags()
	if *flags.ConfigPath == "" {
		readDefaultConfig()
	} else {
		readCustomConfig(*flags.ConfigPath)
	}

	setValues()

	log.Debug("config successfull")
}

func Default() {
	readDefaultConfig()
	setValues()
}

func readDefaultConfig() {
	if err := yaml.Unmarshal(defaultConfig, &config); err != nil {
		log.Fatalf("failed to read default config file: %s", err)
	}
}

func readCustomConfig(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal("failed to open custom config file")
	}
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("failed to read custom config file: %s", err)
	}
}

func setValues() {
	Database = config.Database
	Server = config.Server
}
