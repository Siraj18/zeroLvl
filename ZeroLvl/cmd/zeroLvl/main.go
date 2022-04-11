package main

import (
	"os"
	"time"

	"github.com/siraj18/zeroLvl/internal/handlers"
	"github.com/siraj18/zeroLvl/internal/repositories/modelsrepo"
	"github.com/siraj18/zeroLvl/internal/services/modelsrv"
	"github.com/siraj18/zeroLvl/pkg/inmemory"
	"github.com/siraj18/zeroLvl/pkg/postgres"
	"github.com/siraj18/zeroLvl/pkg/server"
	"github.com/siraj18/zeroLvl/pkg/stancon"
	"github.com/sirupsen/logrus"
)

func main() {
	//conStr := "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable"
	conStr := os.Getenv("connection_string_postgres")
	clusterId := os.Getenv("cluster_id")
	clientId := os.Getenv("client_id")
	channel := os.Getenv("channel")
	durableName := os.Getenv("durrable_name")
	natsUrl := os.Getenv("nats_url")
	address := os.Getenv("address")

	cache := inmemory.NewCache()
	cacheRep := modelsrepo.NewCacheRepository(cache)

	db, err := postgres.NewDb(conStr, 10)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	postgreRep, err := modelsrepo.NewPostgreRepository(db)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	modelsService := modelsrv.NewModelsService(postgreRep, cacheRep)
	if err := modelsService.InitCacheFromDb(); err != nil {
		logrus.Error(err)
	}

	handler := handlers.NewHandler(modelsService)

	stanCon := stancon.NewStanConnection(clusterId, clientId, natsUrl)
	stanHandler := handlers.NewStanHandler(*stanCon, modelsService)
	stanHandler.Subscribe(channel, durableName)

	srv := server.NewServer(address, handler.InitRoutes(), time.Second*3)

	if err := srv.Run(); err != nil {
		logrus.Fatal(err)
	}
}
