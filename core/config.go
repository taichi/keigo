package core

import (
	"github.com/BurntSushi/toml"
	"time"
)

type KeigoConfig struct {
	Address  string        `toml:"keiko_address"`
	Terminal byte          `toml:"terminal_code"`
	Timeout  time.Duration `toml:"timeout"`
	Retry    int           `toml:"retry"`
}

func NewConfig() *KeigoConfig {
	return &KeigoConfig{
		Terminal: '\r',
		Timeout:  60 * time.Second,
		Retry:    3,
	}
}

func LoadConfig(path string) (*KeigoConfig, error) {
	config := NewConfig()
	if _, err := toml.DecodeFile(path, config); err == nil {
		return config, nil
	} else {
		return nil, err
	}
}
