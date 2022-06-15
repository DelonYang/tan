package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var Conf *TomlConfig

type TomlConfig struct {
	Common commonConfig
	Mysql  mysqlConfig
	Redis  redisConfig
}

type commonConfig struct {
	Port int `toml:"port"`
}

type mysqlConfig struct {
	Url    string `toml:"url"`
	Auth   string `toml:"auth"`
	Dbname string `toml:"dbname"`
}

type redisConfig struct {
	Url  string `toml:"url"`
	Auth string `toml:"auth"`
}

func (c *TomlConfig) LoadConfig(env string) {
	if env == "" {
		env = "dev"
	}

	filePath := "../config/config-" + env + ".toml"

	if _, err := os.Stat(filePath); err != nil {
		panic(err)
	}

	if _, err := toml.DecodeFile(filePath, &c); err != nil {
		panic(err)
	}
}

func GetConfEnv() string {
	env := os.Getenv("Tan_ENV")
	if env == "" || env == "dev" {
		env = "dev"
	} else {
		env = os.Args[1]
	}

	return env
}

func init() {
	Conf = new(TomlConfig)
	env := GetConfEnv()
	Conf.LoadConfig(env)
	printLog()
}

func printLog() {
	fmt.Printf("Common: %+v\n", Conf.Common)
	fmt.Printf("mysql: %+v\n", Conf.Mysql)
	fmt.Printf("redis: %+v\n", Conf.Redis)
}
