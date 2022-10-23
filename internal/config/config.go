package config

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"gopkg.in/yaml.v3"
)

const (
	dbPassEscSeq = "{password}"
	password     = "notes_pass"
)

type Grpc struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	Host                 string `yaml:"host"`
	Port                 string `yaml:"port"`
	User                 string `yaml:"user"`
	Name                 string `yaml:"database"`
	Ssl                  string `yaml:"ssl"`
	MaxOpenedConnections int32  `yaml:"max_opened_connections"`
}

type Http struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	Grpc     Grpc     `yaml:"grpc"`
	Database Database `yaml:"database"`
	Http     Http     `yaml:"http"`
}

func Read(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) GetDBConfig() (*pgxpool.Config, error) {
	DbDsn := fmt.Sprintf("host=%s port=%s user=%s password={password} dbname=%s sslmode=%s", c.Database.Host, c.Database.Port, c.Database.User, c.Database.Name, c.Database.Ssl)
	DbDsn = strings.ReplaceAll(DbDsn, dbPassEscSeq, password)

	poolConfig, err := pgxpool.ParseConfig(DbDsn)
	if err != nil {
		return nil, err
	}
	poolConfig.ConnConfig.BuildStatementCache = nil
	poolConfig.ConnConfig.ConnectTimeout = time.Second
	poolConfig.ConnConfig.PreferSimpleProtocol = true
	poolConfig.MaxConns = c.Database.MaxOpenedConnections

	return poolConfig, nil
}

func (g *Grpc) GetAddress() string {
	return net.JoinHostPort(g.Host, g.Port)
}

func (h *Http) GetAddress() string {
	return net.JoinHostPort(h.Host, h.Port)
}
