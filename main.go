package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/kevin07696/produce-service/adapters"
	"github.com/kevin07696/produce-service/domain"
	"github.com/kevin07696/produce-service/generated"
	"github.com/kevin07696/produce-service/handlers"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	var (
		repository       domain.ProductRepositor
		productService   handlers.ProductServicer
		inventoryService domain.InventoryServicer
		productHandler   handlers.ProductHandler
		healthHandler    handlers.HealthHandler
		server           *handlers.Server
	)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to the in-memory database.", "error", err)
		os.Exit(1)
	}

	repository = domain.NewProductRepository(db)
	repository.Migrate()

	inventoryService = adapters.InventoryService{}
	productService = domain.NewProductService(repository, inventoryService)

	productHandler = handlers.NewProductHandler(productService, repository)
	healthHandler = handlers.NewHealthHandler(30)

	server = handlers.NewServer()

	generated.RegisterProductServer(server.Server(), productHandler)
	health.RegisterHealthServer(server.Server(), &healthHandler)

	reflection.Register(server.Server())

	repository.WriteProducts()

	go server.Serve()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Wait for a termination signal
	sig := <-quit
	logger.Info("Received signal. Shutting down...", "signal", sig)

	server.Server().GracefulStop()
}
