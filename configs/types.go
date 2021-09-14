package configs

type Config struct {
	Database *DatabaseConfig `json:"database"`
	Server   *ServerConfig   `json:"server"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ServerConfig struct {
	Address string `json:"address"`
}
