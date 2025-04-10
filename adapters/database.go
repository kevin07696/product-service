package adapters

import (
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitInMemoryDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to the in-memory database.", "error", err)
		os.Exit(1)
	}

	return db
}

// InitDatabase initializes a connection to the PostgreSQL database.
func InitDatabase() *gorm.DB {
	// Construct the DSN (Data Source Name) from environment variables
	dsn := "host=postgres user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_DATABASE") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"

	// Open a connection to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to the PostgreSQL database.", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully connected to the PostgreSQL database.")
	return db
}
