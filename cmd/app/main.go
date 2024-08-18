package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"pxgen.io/user/internal/config"
	"pxgen.io/user/internal/constants"
	"pxgen.io/user/internal/handler"
	"pxgen.io/user/internal/repo"
	"pxgen.io/user/internal/router"
	"pxgen.io/user/internal/utils/log"
)

func main() {

	fmt.Print(constants.BANNER)

	config.Init()

	db := ConnectMySQL()
	defer db.Close()

	userHandler := handler.NewUserHandler(repo.NewUserRepository(db))
	authHandler := handler.NewAuthHandler(repo.NewAuthRepo(db))
	router := router.NewRouter(*userHandler, *authHandler)

	log.Infof("Starting application on port %s", config.App.Port)
	http.ListenAndServe(":"+config.App.Port, router.SetupRouter())

}

func ConnectMySQL() *sql.DB {

	log.Infof("Connecting to MYSQL at %s:%s as %s", config.MySQL.Url, config.MySQL.Port, config.MySQL.User)

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
		log.Logger().Fatal("failed to establish connection with database")
	}

	return db
}
