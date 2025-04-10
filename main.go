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
)

func main() {

	var (
		repository          *domain.ProductRepository
		service             *domain.ProductService
		readProductHandler  handlers.ReadProductHandler
		writeProductHandler handlers.WriteProductHandler
		healthHandler       handlers.HealthHandler
		server              *handlers.Server
	)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	db := adapters.InitDatabase()

	repository = domain.NewProductRepository(db)
	repository.Migrate()

	service = domain.NewProductService(repository)

	readProductHandler = handlers.NewReadProductHandler(repository)
	writeProductHandler = handlers.NewWriteProductHandler(service)
	healthHandler = handlers.NewHealthHandler(30)

	server = handlers.NewServer()

	generated.RegisterProductReadServiceServer(server.Server(), readProductHandler)
	generated.RegisterProductWriteServiceServer(server.Server(), writeProductHandler)
	health.RegisterHealthServer(server.Server(), &healthHandler)

	reflection.Register(server.Server())

	go server.Serve()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	sig := <-quit
	logger.Info("Received signal. Shutting down...", "signal", sig)

	server.Server().GracefulStop()
}
