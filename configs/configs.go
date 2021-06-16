package configs

type Config struct {
	Server *ServerConfig `json:"server"`
	Metric *MetricConfig `json:"metric"`
	DB     *DBConfig     `json:"db"`
	Consul *ConsulConfig `json:"consul"`
	Vault  *VaultConfig  `json:"vault"`
	Proxy  *ProxyConfig  `json:"proxy"`
}

func NewConfig() *Config {
	return &Config{
		Server: &ServerConfig{},
		Metric: &MetricConfig{},
		DB:     &DBConfig{},
		Consul: &ConsulConfig{},
		Vault:  &VaultConfig{},
		Proxy:  &ProxyConfig{},
	}
}

type ProxyConfig struct {
	Addr string `json:"addr"`
}

type ConsulConfig struct {
	Addr  string `env:"CONSUL_ADDR" json:"addr" default:"127.0.0.1:8500"`
	Token string `env:"CONSUL_TOKEN" json:"token"`
	Path  string `env:"CONSUL_PATH" json:"-" default:"service-platform/apigw"`
}

type VaultConfig struct {
	Addr  string `env:"VAULT_ADDR" json:"addr" default:"127.0.0.1:8200"`
	Token string `env:"VAULT_TOKEN" json:"-"`
	Path  string `env:"VAULT_PATH" json:"-" default:"service-platform/apigw"`
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
