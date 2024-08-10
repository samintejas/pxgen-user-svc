package config

import (
	"os"
	"strconv"
	"sync"

	"github.com/rs/zerolog/log"
)

var loaded sync.Once
var App *AppConfigs
var MySQL *MySQLConfigs

func Init() {
	loaded.Do(func() {
		loadConfig()
	})
}

type AppConfigs struct {
	Port string
}

type MySQLConfigs struct {
	Url             string
	Port            string
	Database        string
	User            string
	Password        string
	Network         string
	AllowNativePass bool
	ParseTime       bool
}

func loadConfig() {

	App = &AppConfigs{
		Port: loadString("PXGEN_USR_APP_PORT", "8080"),
	}

	MySQL = &MySQLConfigs{
		Url:             loadString("PXGEN_USR_MYSQL_URL", "127.0.0.1"),
		Port:            loadString("PXGEN_USR_MYSQL_PORT", "3306"),
		Database:        loadString("PXGEN_USR_MYSQL_DATABASE", "pxdt_menu"),
		User:            loadString("PXGEN_USR_MYSQL_USER", "samin"),
		Password:        loadString("PXGEN_USR_MYSQL_PASS", "1"),
		Network:         loadString("PXGEN_USR_MYSQL_NET", "tcp"),
		AllowNativePass: loadBool("PXGEN_USR_MYSQL_ANP", true),
		ParseTime:       loadBool("PXGEN_USR_MYSQL_PT", true),
	}
}

func loadString(key, def string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Info().
			Str("env", key).
			Str("default", def).
			Msg("Config not set, using default value")
		return def
	}
	return value
}

func loadBool(key string, def bool) bool {
	value := os.Getenv(key)
	if value == "" {
		log.Info().
			Str("env", key).
			Bool("default", def).
			Msg("Config not set, using default value")
		return def
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		log.Fatal().
			Str("env", key).
			Err(err).
			Msg("Invalid boolean value")
	}
	return boolValue
}
