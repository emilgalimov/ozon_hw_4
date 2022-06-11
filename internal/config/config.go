package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"database_name"`
}

type Kafka struct {
	Brokers       []string `yaml:"brokers"`
	ConfirmOrders string   `yaml:"confirm_orders"`
	WriteOff      string   `yaml:"write_off"`
	Rejected      string   `yaml:"rejected"`
	Delivered     string   `yaml:"delivered"`
}

type Config struct {
	Kafka   Kafka `yaml:"kafka"`
	Storage struct {
		Db Postgres `yaml:"db"`
	} `yaml:"storage"`
	Pays struct {
		Db Postgres `yaml:"db"`
	} `yaml:"pays"`
	Notifications struct {
		Db Postgres `yaml:"db"`
	} `yaml:"notifications"`
}

func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
