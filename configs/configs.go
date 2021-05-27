package configs

type Config struct {
	Server *ServerConfig `json:"server"`
	Metric *MetricConfig `json:"metric"`
	DB     *DBConfig     `json:"db"`
}

func NewConfig() *Config {
	return &Config{
		Server: &ServerConfig{},
		Metric: &MetricConfig{},
		DB:     &DBConfig{},
	}
}

type DBConfig struct {
	Name string `json:"name" env:"DB_NAME"`
	URL  string `json:"url" env:"DB_URL"`
}

type MetricConfig struct {
	Addr string `json:"addr" env:"METRIC_ADDR" default:":8080"`
}

type ServerConfig struct {
	Addr    string `json:"addr" env:"SERVER_ADDR" default:":9090"`
	Name    string `json:"name"`
	Version string `json:"version" env:"VERSION"`
}
