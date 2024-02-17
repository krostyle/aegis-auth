package configuration

import (
	"fmt"
	"os"

	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {

	env := os.Getenv("ENV")
	var dataSourceName string
	if env == "development" {
		host := os.Getenv("POSTGRES_HOST")
		port := os.Getenv("POSTGRES_PORT")
		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		dbname := os.Getenv("POSTGRES_DB")
		sslmode := os.Getenv("POSTGRES_SSLMODE")
		dataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, user, password, dbname, sslmode)
	} else {
		dataSourceName = os.Getenv("DATABASE_URL")
	}

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening datanase: %w", err)
	}

	if err := migration.RunMigration(db); err != nil {
		return nil, fmt.Errorf("error running migration: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting sql.DB from GORM: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}
