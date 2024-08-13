package config

import (
	"os"
	"strconv"
	"sync"

	"pxgen.io/user/internal/utils/log"
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
	Port           string
	PerfomanceMode bool
	GodMode        bool
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
		Port:           loadString("PXGEN_USR_APP_PORT", "8080"),
		PerfomanceMode: loadBool("PXGEN_USR_APP_PERF_MODE", false),
		GodMode:        loadBool("PXGEN_USR_APP_GOD_MODE", false),
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
		log.Infof("env variable not found , using default, KEY=%s, VALUE=%s", key, def)
		return def
	}
	return value
}

func loadBool(key string, def bool) bool {
	value := os.Getenv(key)
	if value == "" {
		log.Infof("env variable not found , using default, KEY=%s, VALUE=%s", key, strconv.FormatBool(def))
		return def
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		log.Error("invalid boolean value for the configuration")
	}
	return boolValue
}
