package main

import (
	"fmt"
	"go-covid/api"
	"go-covid/app/modules"
	"go-covid/config"
	"go-covid/utils"
	"log"
	"net/http"
	"os"

	_ "go-covid/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title API Jasa Pengiriman
// @version 1.0
// @description Berikut API Jasa Pengiriman
// @host 18.139.225.163:8080
// @BasePath /
func main() {
	config := config.GetConfig()

	dbCon := utils.NewConnectionDatabase(config)

	defer dbCon.CloseConnection()

	controllers := modules.RegistrationModules(dbCon, config)

	e := echo.New()
	handleSwagger := echoSwagger.WrapHandler
	e.GET("/swagger/*", handleSwagger)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	})
	api.RegistrationPath(e, controllers)

	go func() {
		address := fmt.Sprintf(":%d", config.App.Port)
		if err := e.Start(address); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	<-quit
}
