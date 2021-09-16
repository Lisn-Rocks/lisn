package configs

type Config struct {
	Database *DatabaseConfig `yaml:"database"`
	Server   *ServerConfig   `yaml:"server"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type ServerConfig struct {
	Address string `yaml:"address"`
}
