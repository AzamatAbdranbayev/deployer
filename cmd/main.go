package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	configs "github.com/AzamatAbdranbayev/deployer/internal/config"
	"github.com/AzamatAbdranbayev/deployer/internal/handler"
	service2 "github.com/AzamatAbdranbayev/deployer/internal/service"
	"github.com/AzamatAbdranbayev/deployer/internal/transport/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var err error
	var errCh chan error

	file, errFile := os.Open("internal/config/config.json")
	if errFile != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	cfg := configs.NewConfig("server_name", "1.0.0")
	errCfg := decoder.Decode(&cfg)
	if errCfg != nil {
		log.Fatal(err)
	}

	go graceFullyShutdown(errCh)
	if err != nil {
		log.Fatal(err.Error())
	}
	staticService, _ := service2.NewManager(cfg, service2.CreateHttpClient())
	staticHandler := handler.NewManager(staticService)
	HTTPServer := http.NewServer(cfg, staticHandler)
	err = HTTPServer.StartHTTPServer(ctx)
	if err != nil {
		return err
	}
	select {
	case err = <-errCh:
		return err
	}
}

func graceFullyShutdown(errorCh chan error) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	errorCh <- fmt.Errorf("%s", <-ch)
	return
}
