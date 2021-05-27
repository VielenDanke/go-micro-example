package configs

type Config struct {
	Server *ServerConfig `json:"server"`
	Metric *MetricConfig `json:"metric"`
}

func NewConfig() *Config {
	return &Config{
		Server: &ServerConfig{},
		Metric: &MetricConfig{},
	}
}

type MetricConfig struct {
	Addr string `json:"addr" env:"METRIC_ADDR" default:":8080"`
}

type ServerConfig struct {
	Addr    string `json:"addr" env:"SERVER_ADDR" default:":9090"`
	Name    string `json:"name"`
	Version string `json:"version" env:"VERSION"`
}
