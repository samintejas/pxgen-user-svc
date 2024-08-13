package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"pxgen.io/user/internal/config"
	"pxgen.io/user/internal/constants"
	"pxgen.io/user/internal/handler"
	"pxgen.io/user/internal/repo"
	"pxgen.io/user/internal/router"
)

func main() {

	fmt.Print(constants.BANNER + "\n")

	configLogging()
	log.Info().Msg("Starting pxgen user management service ..")
	config.Init()
	db := ConnectMySQL()
	defer db.Close()

	userHandler := handler.NewUserHandler(repo.NewUserRepository(db))
	router := router.NewRouter(*userHandler)

	log.Info().Str("Port", config.App.Port).Msg("starting server")
	http.ListenAndServe(":"+config.App.Port, router.SetupRouter())

}

func ConnectMySQL() *sql.DB {

	cfg := mysql.Config{
		User:                 config.MySQL.User,
		Passwd:               config.MySQL.Password,
		Addr:                 config.MySQL.Url + ":" + config.MySQL.Port,
		DBName:               config.MySQL.Database,
		Net:                  config.MySQL.Network,
		AllowNativePasswords: config.MySQL.AllowNativePass,
		ParseTime:            config.MySQL.ParseTime,
	}

	db, _ := sql.Open("mysql", cfg.FormatDSN())

	err := db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open mysql connection")
	}

	return db
}

func configLogging() {
	perfMode, ok := os.LookupEnv("PXGEN_USR_APP_PERF_MODE")
	if !ok {
		perfMode = "false"
	}
	if perfMode == "false" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: zerolog.TimeFieldFormat})
	} else {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	}
}
