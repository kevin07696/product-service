package adapters

import (
	"log/slog"
	"os"

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
