package config

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`

	Database struct {
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		DBName    string `yaml:"dbname"`
		Charset   string `yaml:"charset"`
		ParseTime bool   `yaml:"parseTime"`
		Loc       string `yaml:"loc"`
	} `yaml:"database"`
}

var GlobalConfig Config

func Init() error {
	// 读取配置文件
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return fmt.Errorf("read config file error: %v", err)
	}

	// 解析配置
	if err := yaml.Unmarshal(data, &GlobalConfig); err != nil {
		return fmt.Errorf("parse config error: %v", err)
	}

	// 从环境变量覆盖配置
	if dbPass := os.Getenv("DB_PASSWORD"); dbPass != "" {
		GlobalConfig.Database.Password = dbPass
	}
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		GlobalConfig.Database.Host = dbHost
	}
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		if port, err := strconv.Atoi(dbPort); err == nil {
			GlobalConfig.Database.Port = port
		}
	}

	return nil
}
