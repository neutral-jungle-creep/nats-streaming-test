package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"nats-listener/configs"
	"nats-listener/internal/caching"
	"nats-listener/internal/delivery"
	pkg "nats-listener/internal/delivery/http"
	"nats-listener/internal/delivery/nats"
	"nats-listener/internal/service"
	"nats-listener/internal/storage"
	"nats-listener/internal/storage/pgSQL"
	"nats-listener/pkg/logger"
	"net/http"
)

func main() {
	log := logger.NewLogger()
	log.Info("logger has been started")

	config := configs.NewConfig()
	log.Info("config has been init")

	dbConn, err := pgSQL.NewPgConnect(config.DataBase)
	if err != nil {
		log.Fatalf("can't connect to data base by link: %s", err.Error())
	}
	log.Infof("connect to data base is successful by link: %s", config.DataBase.ConnLink)
	defer dbConn.Close(context.Background())

	//todo переделать позже
	cache := caching.NewCache()
	if err := cache.FillCache(dbConn); err != nil {
		log.Errorf("can't fill cache: %s", err.Error())
	}
	log.Info("cache filled successfully")

	stor := storage.NewStorage(dbConn, cache, log)
	serv := service.NewService(stor)
	handler := delivery.NewHandler(serv)
	var valid *validator.Validate
	valid = validator.New()

	natsQueue, err := nats.NewNatsQueue(config.Nats, valid)
	if err != nil {
		log.Fatalf("can't connect to nats: %s", err.Error())
	}
	log.Infof("connect to nats is successfully: url = [%s], cluster = [%s], client = [%s]",
		config.Nats.URL, config.Nats.ClusterID, config.Nats.ClientID)
	var messageQueue nats.OrdersMessageQueue = natsQueue
	defer messageQueue.Close()

	messageQueue.OnNewMessage(handler.HandleNewOrder)

	go func() {
		server := new(pkg.Server)

		log.Infof("server started on: http://localhost:%s", config.Http.Port)

		http.HandleFunc("/get-order", handler.GetOrderById)

		fs := http.FileServer(http.Dir("./frontend"))
		http.Handle("/", fs)

		log.Fatal(server.Run(config.Http.Port, nil))
	}()

	ch := make(chan int)
	<-ch
}
