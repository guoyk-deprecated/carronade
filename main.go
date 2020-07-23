package main

import (
	"context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	optBind       = ":9000"
	optDataDriver = "mysql"
	optDataSource = "root@tcp(127.0.0.1:3306)/carronade?charset=utf8mb4&parseTime=true"
)

func envStr(out *string, key string) {
	if val := strings.TrimSpace(os.Getenv(key)); val != "" {
		*out = val
	}
}

func exit(err *error) {
	if *err != nil {
		log.Printf("exited with error: %s", (*err).Error())
		os.Exit(1)
	} else {
		log.Println("exited")
	}
}

func init() {
	envStr(&optBind, "CARRONADE_BIND")
	envStr(&optDataDriver, "CARRONADE_DATA_DRIVER")
	envStr(&optDataSource, "CARRONADE_DATA_SOURCE")
}

func main() {
	var err error
	defer exit(&err)

	// connect database
	var db *gorm.DB
	if db, err = gorm.Open(optDataDriver, optDataSource); err != nil {
		return
	}
	defer db.Close()

	// create server
	e := echo.New()
	e.HidePort = true
	e.HideBanner = true
	e.Use(middleware.Recover())
	defer e.Shutdown(context.Background())

	log.Printf("listen at %s", optBind)

	chErr := make(chan error, 1)
	chSig := make(chan os.Signal, 1)

	// start server and wait signals
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		chErr <- e.Start(optBind)
	}()

	select {
	case err = <-chErr:
	case sig := <-chSig:
		log.Printf("signal caught: %s", sig.String())
	}
}
