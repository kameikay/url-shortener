package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kameikay/url-shortener/configs"
	"github.com/kameikay/url-shortener/internal/infra/repository"
	"github.com/kameikay/url-shortener/internal/infra/web/controller"
	"github.com/kameikay/url-shortener/internal/infra/web/handlers"
	"github.com/kameikay/url-shortener/internal/infra/web/webserver"

	_ "github.com/lib/pq"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	dbConnection, err := sql.Open(configs.DBDriver, fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	dbConnection.SetConnMaxLifetime(time.Minute * 3)
	defer dbConnection.Close()

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webserver.MountMiddlewares()

	repo := repository.NewRepository(dbConnection)
	handler := handlers.NewHandler(repo)
	controller := controller.NewController(webserver.Router, handler)
	controller.Route()

	webserver.Start()
}
