package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type ServerConfig struct {
	Engine  EngineConfig  `yaml:"engine"`
	Network NetworkConfig `yaml:"network"`
	Logging LoggingConfig `yaml:"logging"`
}

type EngineConfig struct {
	Type string `yaml:"type"`
}

type NetworkConfig struct {
	Address        string `yaml:"address"`
	MaxConnections int    `yaml:"max_connections"`
	MaxMessageSize string `yaml:"max_message_size"`
	IdleTimeout    string `yaml:"idle_timeout"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Output string `yaml:"output"`
}

const (
	defaultEngineType = "in_memory"

	defaultAddress        = "127.0.0.1:3223"
	defaultMaxConnections = 100
	defaultMaxMessageSize = "4KB"
	defaultIdleTimeout    = "5m"

	defaultLoggingLevel  = "debug"
	defaultLoggingOutput = "log/output.log"
)

func initiateDefaultConfig() ServerConfig {
	return ServerConfig{
		Engine: EngineConfig{
			Type: defaultEngineType,
		},
		Network: NetworkConfig{
			Address:        defaultAddress,
			MaxConnections: defaultMaxConnections,
			MaxMessageSize: defaultMaxMessageSize,
			IdleTimeout:    defaultIdleTimeout,
		},
		Logging: LoggingConfig{
			Level:  defaultLoggingLevel,
			Output: defaultLoggingOutput,
		},
	}
}

func NewConfigFromBytes(data []byte) ServerConfig {
	cfg := initiateDefaultConfig()

	err := yaml.Unmarshal(data, &cfg)

	if err != nil {
		fmt.Printf("Cannot unarshal config: [%v]", err)
		return cfg
	}

	return cfg
}

func NewConfigFromFile(path string) ServerConfig {
	data, err := os.ReadFile(filepath.Clean(path))

	if err != nil {
		fmt.Printf("Cannot read config from file. Return default: %s [%v]", path, err)

		return initiateDefaultConfig()
	}

	return NewConfigFromBytes(data)
}
